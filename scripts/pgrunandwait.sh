#! /bin/bash

APP_DB_USER=cloudbilling
APP_DB_PWD=cloudbilling
docker run --rm -d -p 5432:5432 -v cloudbilling-pgdata:/var/lib/postgresql/data --health-cmd pg_isready --health-interval 0.5s -e POSTGRES_USER=${APP_DB_USER} --name cloudbilling-postgres postgres:10-alpine
until [ `docker inspect --format "{{json .State.Health.Status }}" cloudbilling-postgres` = '"healthy"' ];
do 
    echo "waiting for postgres container..."
    sleep 0.5
done
echo "pg ready"