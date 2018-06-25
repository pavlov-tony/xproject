FROM golang:1.10.2-alpine3.7

WORKDIR /go/src/github.com/pavlov-tony/xproject

VOLUME .:/go/src/github.com/pavlov-tony/xproject

# ENV APP_DB_USERNAME docker
# ENV APP_DB_PASSWORD docker
# ENV APP_DB_NAME docker

RUN apk update
RUN apk add git
RUN go get -u github.com/golang/dep/cmd/dep

COPY . .

EXPOSE 8080
