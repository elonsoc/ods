--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2 (Debian 15.2-1.pgdg110+1)

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

--
-- Name: affiliation; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.affiliation AS ENUM (
    'facstaff',
    'alumni',
    'student'
);


ALTER TYPE public.affiliation OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: app_owner; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.app_owner (
    app_id uuid NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE public.app_owner OWNER TO postgres;

--
-- Name: applications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.applications (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    api_key text DEFAULT gen_random_uuid() NOT NULL,
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
    email text NOT NULL,
    affiliation public.affiliation NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: app_owner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.app_owner (app_id, user_id) FROM stdin;
\.


--
-- Data for Name: applications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.applications (id, name, description, api_key, is_valid) FROM stdin;
\.


--
-- Data for Name: elon_ods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.elon_ods (elon_id, ods_id) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, first_name, last_name, email, affiliation) FROM stdin;
\.


--
-- Name: app_owner app_owner_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.app_owner
    ADD CONSTRAINT app_owner_pkey PRIMARY KEY (app_id, user_id);


--
-- Name: applications applications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_pkey PRIMARY KEY (id);


--
-- Name: elon_ods elon_ods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.elon_ods
    ADD CONSTRAINT elon_ods_pkey PRIMARY KEY (elon_id, ods_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: api_key_is_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX api_key_is_unique ON public.applications USING btree (api_key);


--
-- Name: elon_id_is_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX elon_id_is_unique ON public.elon_ods USING btree (elon_id);


--
-- Name: ods_id_is_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ods_id_is_unique ON public.elon_ods USING btree (ods_id);


--
-- Name: app_owner app_owner_app_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.app_owner
    ADD CONSTRAINT app_owner_app_id_fkey FOREIGN KEY (app_id) REFERENCES public.applications(id);


--
-- Name: app_owner app_owner_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.app_owner
    ADD CONSTRAINT app_owner_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: elon_ods elon_ods_ods_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.elon_ods
    ADD CONSTRAINT elon_ods_ods_id_fkey FOREIGN KEY (ods_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--
