--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Homebrew)
-- Dumped by pg_dump version 15.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: applications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.applications (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    owners text NOT NULL,
    api_key text NOT NULL,
    is_valid boolean NOT NULL
);


ALTER TABLE public.applications OWNER TO postgres;

--
-- Name: elon_ods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.elon_ods (
    elon_id text NOT NULL,
    ods_id uuid NOT NULL
);


ALTER TABLE public.elon_ods OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: applications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.applications (id, name, description, owners, api_key, is_valid) FROM stdin;
\.


--
-- Data for Name: elon_ods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.elon_ods (elon_id, ods_id) FROM stdin;
\.

--
-- Name: applications applications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: elon_ods elon_ods_ods_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.elon_ods
    ADD CONSTRAINT elon_ods_ods_id_fkey FOREIGN KEY (ods_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--
