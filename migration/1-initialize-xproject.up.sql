SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

CREATE SCHEMA xproject;
ALTER SCHEMA xproject OWNER TO postgres;

SET default_tablespace = '';
SET default_with_oids = false;

CREATE TABLE xproject.accounts (
    id integer NOT NULL,
    gcp_account_info text
);
ALTER TABLE xproject.accounts OWNER TO postgres;

CREATE SEQUENCE xproject.accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE xproject.accounts_id_seq OWNER TO postgres;
ALTER SEQUENCE xproject.accounts_id_seq OWNED BY xproject.accounts.id;

CREATE TABLE xproject.gcp_csv_files (
    id integer NOT NULL,
    name text,
    bucket text,
    time_created timestamp without time zone,
    account_id integer
);
ALTER TABLE xproject.gcp_csv_files OWNER TO postgres;

CREATE SEQUENCE xproject.gcp_csv_files_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE xproject.gcp_csv_files_id_seq OWNER TO postgres;
ALTER SEQUENCE xproject.gcp_csv_files_id_seq OWNED BY xproject.gcp_csv_files.id;

CREATE TABLE xproject.service_bills (
    id integer NOT NULL,
    line_item text,
    start_time timestamp without time zone,
    end_time timestamp without time zone,
    cost real,
    currency text,
    project_id text,
    description text,
    gcp_csv_file_id integer
);
ALTER TABLE xproject.service_bills OWNER TO postgres;

CREATE SEQUENCE xproject.service_bills_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE xproject.service_bills_id_seq OWNER TO postgres;
ALTER SEQUENCE xproject.service_bills_id_seq OWNED BY xproject.service_bills.id;

ALTER TABLE ONLY xproject.accounts ALTER COLUMN id SET DEFAULT nextval('xproject.accounts_id_seq'::regclass);
ALTER TABLE ONLY xproject.gcp_csv_files ALTER COLUMN id SET DEFAULT nextval('xproject.gcp_csv_files_id_seq'::regclass);
ALTER TABLE ONLY xproject.service_bills ALTER COLUMN id SET DEFAULT nextval('xproject.service_bills_id_seq'::regclass);

ALTER TABLE ONLY xproject.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);
ALTER TABLE ONLY xproject.gcp_csv_files
    ADD CONSTRAINT gcp_csv_files_pkey PRIMARY KEY (id);
ALTER TABLE ONLY xproject.service_bills
    ADD CONSTRAINT service_bills_pkey PRIMARY KEY (id);

ALTER TABLE ONLY xproject.gcp_csv_files
    ADD CONSTRAINT gcp_csv_files_fkey FOREIGN KEY (account_id) REFERENCES xproject.accounts(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY xproject.service_bills
    ADD CONSTRAINT service_bills_fkey FOREIGN KEY (gcp_csv_file_id) REFERENCES xproject.gcp_csv_files(id) ON UPDATE CASCADE ON DELETE CASCADE;
