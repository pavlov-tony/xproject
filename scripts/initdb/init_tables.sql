CREATE SCHEMA xproject;

CREATE TABLE xproject.accounts (
    id SERIAL PRIMARY KEY,
    gcp_account_info text
);

CREATE TABLE xproject.gcp_csv_files (
    id SERIAL PRIMARY KEY,
    name text,
    bucket text,
    time_created timestamp without time zone,
    account_id integer REFERENCES xproject.accounts(id)
);

CREATE TABLE xproject.service_bills (
    id SERIAL PRIMARY KEY,
    line_item text,
    start_time timestamp without time zone,
    end_time timestamp without time zone,
    cost real,
    currency text,
    project_id text,
    description text,
    gcp_csv_file_id integer REFERENCES xproject.gcp_csv_files(id)
);
