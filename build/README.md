# Docker Guide

## To install go app:
### Building and running go container
```shell
$ cd xproject/
$ docker build -t myimage -f $(pwd)/build/Golang.Dockerfile .
$ docker run -it --rm --name myapp -v $(pwd):/go/src/github.com/pavlov-tony/xproject -p 8080:8080 myimage sh
```

### Building and running postgresql container
1. [Write init.sql script for container](https://stackoverflow.com/questions/26598738/how-to-create-user-database-in-script-for-docker-postgres)
```shell
$ cd xproject/build
$ docker build -t pgimage -f $(pwd)/PG.Dockerfile .
$ docker run --rm --name dbapp -p 5432:5432 pgimage
```
Exec sql script in container:
```shell
$ docker exec -it dbapp sh
$ psql -h localhost -U docker -d docker -a -f initdb/create_tables.sql
```

### Remove ALL containers, images and volumes:
```shell
$ docker rm $(docker ps -a -q)
$ docker rmi $(docker images -q)
$ docker volume rm $(docker volume ls -qf dangling=true)
```
