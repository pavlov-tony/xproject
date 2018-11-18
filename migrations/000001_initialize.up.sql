CREATE SCHEMA cloudbilling;

CREATE TABLE cloudbilling.accounts (
    id SERIAL PRIMARY KEY,
    gcp_account_info text
);

CREATE TABLE cloudbilling.gcp_csv_files (
    id SERIAL PRIMARY KEY,
    name text,
    bucket text,
    time_created timestamp without time zone,
    account_id integer REFERENCES cloudbilling.accounts(id)
);

CREATE TABLE cloudbilling.service_bills (
    id SERIAL PRIMARY KEY,
    line_item text,
    start_time timestamp without time zone,
    end_time timestamp without time zone,
    cost real,
    currency text,
    project_id text,
    description text,
    gcp_csv_file_id integer REFERENCES cloudbilling.gcp_csv_files(id)
);
