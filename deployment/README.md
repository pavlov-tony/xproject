# Docker Compose Guide

# TODO
* web app stops with status "Exit 137" (now idk how to fix)

### Deployment with docker-compose:
```shell
$ cd xproject/
```
up/down/start/stop:
```shell
$ docker-compose -f deployment/docker-compose.yml up/down/start/stop
```
Exec to web service:
```shell
$ docker-compose -f deployment/docker-compose.yml exec golang sh
```
