--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4
-- Dumped by pg_dump version 10.4

-- Started on 2018-07-07 20:52:57

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2803 (class 0 OID 0)
-- Dependencies: 2802
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- TOC entry 8 (class 2615 OID 16579)
-- Name: xproject; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA xproject;


ALTER SCHEMA xproject OWNER TO postgres;

--
-- TOC entry 2 (class 3079 OID 12924)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2805 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 1 (class 3079 OID 16384)
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- TOC entry 2806 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 198 (class 1259 OID 24585)
-- Name: reports; Type: TABLE; Schema: xproject; Owner: postgres
--

CREATE TABLE xproject.reports (
    id integer NOT NULL,
    account_id text,
    line_item text,
    start_time text,
    end_time text,
    cost real,
    currency text,
    project_id text,
    description text
);


ALTER TABLE xproject.reports OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 24593)
-- Name: reports_id_seq; Type: SEQUENCE; Schema: xproject; Owner: postgres
--

CREATE SEQUENCE xproject.reports_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE xproject.reports_id_seq OWNER TO postgres;

--
-- TOC entry 2807 (class 0 OID 0)
-- Dependencies: 199
-- Name: reports_id_seq; Type: SEQUENCE OWNED BY; Schema: xproject; Owner: postgres
--

ALTER SEQUENCE xproject.reports_id_seq OWNED BY xproject.reports.id;


--
-- TOC entry 2673 (class 2604 OID 24596)
-- Name: reports id; Type: DEFAULT; Schema: xproject; Owner: postgres
--

ALTER TABLE ONLY xproject.reports ALTER COLUMN id SET DEFAULT nextval('xproject.reports_id_seq'::regclass);


--
-- TOC entry 2675 (class 2606 OID 24592)
-- Name: reports reports_pkey; Type: CONSTRAINT; Schema: xproject; Owner: postgres
--

ALTER TABLE ONLY xproject.reports
    ADD CONSTRAINT reports_pkey PRIMARY KEY (id);


-- Completed on 2018-07-07 20:52:57

--
-- PostgreSQL database dump complete
--

