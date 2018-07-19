#!/bin/bash
# NOTE: dont use example (git tracked).
# Copy example to yourprefix.setenv.sh (git ignore) file, and edit envs.
# Otherwise, you can add important information by mistake to the repository.

# postgres database env
export APP_DB_PG_HOST=0.0.0.0
export APP_DB_PG_PORT=5432
export APP_DB_PG_NAME=yourDBName
export APP_DB_PG_USER=yourDBUser
export APP_DB_PG_PWD=password

# gcp env
export GOOGLE_APPLICATION_CREDENTIALS=yourCredPath.json
export APP_PROJECT_ID=yourProjId
export APP_PROJECT_BUCKET=yourBucket
export APP_PROJECT_CSV_OBJECT=yourBillingObjInBucket.csv
