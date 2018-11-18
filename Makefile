include build/package/dev_config.env

PROOT=/go/src/github.com/radisvaliullin/xproject
# image name for docker
IMAGE_NAME=xproject

DOCKER_RUN_DEV=@docker run --rm -i -v `pwd`:${PROOT} -v `pwd`/build/.gometalinter.json:/go/src/.gometalinter.json -w ${PROOT}
APP_DB_USER=xproject
APP_DB_PWD=xproject

.PHONY: db/volume
db/volume:
	@docker volume create xproject-pgdata

.PHONY: build-base
build-base:
	@echo "::: building base image"
	@docker build -f build/package/Base.Dockerfile -t $(IMAGE_NAME):base .

.PHONY: install
install: db/volume build-base

.PHONY: db/run
db/run:
	@echo "::: running postgres"
	@scripts/pgrunandwait.sh

#check if a variable is set
.PHONY: check-%
check-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "$* missing"; \
		exit 1; \
	fi

.PHONY: db/migrations
db/migrations: check-CMD db/run
	@echo "::: running migrations..."
	-@migrate -database "postgres://${APP_DB_USER}:${APP_DB_PWD}@localhost:5432/xproject?sslmode=disable" -path migrations $$CMD
	@docker stop xproject-postgres

#RUN actions
.PHONY: unit-test
unit-test:
	@echo "::: running unit tests"
	$(DOCKER_RUN_DEV) $(IMAGE_NAME):base sh -c "go test -v ./..."

.PHONY: lint
lint:
	@echo "::: running code lint"
	$(DOCKER_RUN_DEV) $(IMAGE_NAME):base sh -c "gometalinter --line-length=100 ./... --config=/go/src/.gometalinter.json"

.PHONY: cover
cover:
	@echo "::: generating go coverage report"
	$(DOCKER_RUN_DEV) $(IMAGE_NAME):base sh -c "go test ./... -coverprofile=../c.out && go tool cover -html=../c.out -o ../coverage.html && cat ../coverage.html" > coverage.html
	@xdg-open coverage.html

.PHONY: run
run: check-CMD
	@echo "::: running command inside container"
	$(DOCKER_RUN_DEV) $(IMAGE_NAME):base sh -c "$(CMD)"

.PHONY: deps
deps:
	@echo "::: installing golang dependencies"
	$(DOCKER_RUN_DEV) $(IMAGE_NAME):base sh -c "dep ensure -v -update"

.PHONY: dep-add
dep-add: check-PKG
	@echo "::: installing package $(PKG)"
	$(DOCKER_RUN_DEV) $(IMAGE_NAME):base sh -c "dep ensure -add $(PKG)"

#COMPOSE actions
.PHONY: up
up:
	@echo "::: building dev environment on port 8080"
	@docker-compose -f `pwd`/deployments/docker-compose.yml up

.PHONY: down
down:
	@echo " ::: tear down dev environment"
	@docker-compose -f `pwd`/deployments/docker-compose.yml down

.PHONY: debug
debug:
	@echo "::: debugging dev environment"
	@docker-compose -f `pwd`/deployments/docker-compose-debug.yml up --build

.PHONY: debug-down
debug-down:
	@echo "::: tear down debug dev environment"
	@docker-compose -f `pwd`/deployments/docker-compose-debug.yml down

.PHONY: integration-test
integration-test:
	@echo "::: running integration tests"
	-@docker-compose -f `pwd`/deployments/docker-compose-integration.yml up api
	@docker-compose -f `pwd`/deployments/docker-compose-integration.yml down
