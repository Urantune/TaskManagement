--
-- PostgreSQL database dump
--

\restrict xaPl9NiFcRMiuYarmo2OtZ6zHJbm5tYqurIFWWmKse0DJHTkic3NlcnkmCaXJ0U

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: tasks; Type: TABLE; Schema: public; Owner: urantune
--

CREATE TABLE public.tasks (
    id integer NOT NULL,
    name text,
    status text,
    user_id bigint
);


ALTER TABLE public.tasks OWNER TO urantune;

--
-- Name: tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: urantune
--

CREATE SEQUENCE public.tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tasks_id_seq OWNER TO urantune;

--
-- Name: tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: urantune
--

ALTER SEQUENCE public.tasks_id_seq OWNED BY public.tasks.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: urantune
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email text NOT NULL,
    password_hash text NOT NULL,
    role text DEFAULT 'user'::text NOT NULL
);


ALTER TABLE public.users OWNER TO urantune;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: urantune
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO urantune;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: urantune
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: tasks id; Type: DEFAULT; Schema: public; Owner: urantune
--

ALTER TABLE ONLY public.tasks ALTER COLUMN id SET DEFAULT nextval('public.tasks_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: urantune
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: urantune
--

COPY public.tasks (id, name, status, user_id) FROM stdin;
4	Dc	Ok	1
5	test	Ok	1
6	test	Ok	1
7	test	Ok	1
8	test	Ok	1
9	test	Ok	1
10	test	Ok	1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: urantune
--

COPY public.users (id, created_at, updated_at, deleted_at, email, password_hash, role) FROM stdin;
1	2026-01-06 16:11:05.050847+07	2026-01-06 16:11:05.050847+07	\N	ngon	$2a$10$5ucFGPhd18.qaKj.L4gLles95YqU6sYm2Za2n86K9Um45cx51DCtC	user
2	2026-01-06 21:33:57.215752+07	2026-01-06 21:33:57.215752+07	\N	hehe	$2a$10$CR0sYAHrXPaH6YtD6W.hUuTe8LmCjcRBjn15Deee7xJNH1ayDfIr.	user
3	2026-01-07 08:20:37.416811+07	2026-01-07 08:20:37.416811+07	\N	ok	$2a$10$aizM0G1.YDiaY6SMn4MoL.stQyrs0XtTQyYEPwbbiHq2JllV2aSg.	user
4	2026-01-07 08:21:06.168704+07	2026-01-07 08:21:06.168704+07	\N	homnay	$2a$10$vYjcgBVfMurMDEo9NKIFA.OEQBeHSpRdLZO2vEsQwJs6qbZ.zMzam	user
5	2026-01-06 16:11:05.050847+07	2026-01-06 16:11:05.050847+07	\N	admin	$2a$10$5ucFGPhd18.qaKj.L4gLles95YqU6sYm2Za2n86K9Um45cx51DCtC	admin
\.


--
-- Name: tasks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: urantune
--

SELECT pg_catalog.setval('public.tasks_id_seq', 11, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: urantune
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: urantune
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: urantune
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: urantune
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: urantune
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: tasks fk_tasks_user; Type: FK CONSTRAINT; Schema: public; Owner: urantune
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT fk_tasks_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

\unrestrict xaPl9NiFcRMiuYarmo2OtZ6zHJbm5tYqurIFWWmKse0DJHTkic3NlcnkmCaXJ0U

