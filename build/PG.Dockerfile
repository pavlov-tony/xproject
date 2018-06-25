FROM postgres:10-alpine

# set env
COPY initdb/init.sql /docker-entrypoint-initdb.d/

VOLUME ./initdb:initdb/
COPY initdb/ initdb/
