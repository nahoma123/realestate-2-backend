--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0 (Debian 16.0-1.pgdg120+1)
-- Dumped by pg_dump version 16.0 (Debian 16.0-1.pgdg120+1)

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
-- Name: properties; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.properties (
    id bigint NOT NULL,
    property_id text,
    status text,
    user_id text,
    amount numeric,
    address text,
    coordinate text,
    location text,
    postal_code text,
    street_address text,
    property_type text,
    images character varying(255),
    reception_number bigint,
    bed_number bigint,
    bath_number bigint,
    property_details text,
    epc text,
    is_student_property boolean,
    features character varying(255),
    furnished text,
    next_inspection_date timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    latitude numeric,
    longitude numeric,
    landlord_id text,
    tenant_id text
);


ALTER TABLE public.properties OWNER TO postgres;

--
-- Name: properties_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.properties_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.properties_id_seq OWNER TO postgres;

--
-- Name: properties_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.properties_id_seq OWNED BY public.properties.id;


--
-- Name: real_estates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.real_estates (
    id bigint NOT NULL,
    real_estate_id text,
    email text,
    address text,
    phone_number text,
    why_joined bigint,
    preferred_time timestamp with time zone,
    status text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.real_estates OWNER TO postgres;

--
-- Name: real_estates_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.real_estates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.real_estates_id_seq OWNER TO postgres;

--
-- Name: real_estates_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.real_estates_id_seq OWNED BY public.real_estates.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    user_id text,
    first_name text,
    middle_name text,
    last_name text,
    email text,
    enrollment text,
    phone text,
    password text,
    user_name text,
    gender text,
    status text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    role text,
    reset_code bigint
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: properties id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.properties ALTER COLUMN id SET DEFAULT nextval('public.properties_id_seq'::regclass);


--
-- Name: real_estates id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.real_estates ALTER COLUMN id SET DEFAULT nextval('public.real_estates_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: properties; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.properties (id, property_id, status, user_id, amount, address, coordinate, location, postal_code, street_address, property_type, images, reception_number, bed_number, bath_number, property_details, epc, is_student_property, features, furnished, next_inspection_date, created_at, updated_at, latitude, longitude) FROM stdin;
4	e663ed2c-a137-4d2c-bc5f-2ae35998e7c9	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:30.361118+00	2023-11-01 15:38:30.361118+00	\N	\N
5	7ec1fa82-ae91-4a12-8d3e-6b069c0d649c	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:37.752774+00	2023-11-01 15:38:37.752775+00	\N	\N
6	d0d4561a-555b-408d-92d7-fe286e5d04c0	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:39.005554+00	2023-11-01 15:38:39.005554+00	\N	\N
7	117acc20-b2cb-4249-b876-51e919e2b57d	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:40.650626+00	2023-11-01 15:38:40.650626+00	\N	\N
8	ebc89bca-064d-4e58-9eb9-d29135eb5ced	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:41.663341+00	2023-11-01 15:38:41.663342+00	\N	\N
9	82ed9563-7e40-44c4-b7c4-27add7a8cf12	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:42.736886+00	2023-11-01 15:38:42.736886+00	\N	\N
10	12b28423-a022-47c7-b60c-08a176888784	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:43.818296+00	2023-11-01 15:38:43.818296+00	\N	\N
11	1d9c2ddc-b236-4fc6-a98d-11976bf94d3a	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:44.872802+00	2023-11-01 15:38:44.872802+00	\N	\N
12	b2c69563-a71a-4306-92b8-3cc6d0a38780	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:45.752265+00	2023-11-01 15:38:45.752265+00	\N	\N
13	2f8e1508-1a46-4e75-b7dc-6ca33422ddb6	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:46.611797+00	2023-11-01 15:38:46.611797+00	\N	\N
14	04d4c948-2581-4dfc-bcbb-e64db794115d	Active	00000000-0000-0000-0000-000000000000	4000	134 Main St	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-12-31 23:59:45+00 BC	2023-11-01 15:38:49.425832+00	2023-11-01 15:38:49.425832+00	\N	\N
15	1e5fe0ee-645c-482d-8e3f-f3cc1d0a9519	Active	00000000-0000-0000-0000-000000000000	2000	address 1	\N	\N		\N	Condo	1700996480681550435.png,1700996480686078610.webp	12	5	2	best property out there	C	f	Appliances	Furnished	2023-11-26 10:33:25.878+00	2023-11-26 11:01:20.690632+00	2023-11-26 11:01:20.690632+00	38.8951	-77.0364
16	4a38ea60-54b8-4747-83a6-9c48b0e8d841	Active	00000000-0000-0000-0000-000000000000	20000	address 1	\N	\N		\N	House	1700996732840296262.png,1700996732847316399.webp	12	10	4	best property out there	C	f	Gym	Furnished	2023-12-26 11:04:34+00	2023-11-26 11:05:32.855059+00	2023-11-26 11:05:32.855059+00	38.8951	-77.0364
1	021ab7df-ace5-4997-bbce-6c86dc311695	Cancelled	00000000-0000-0000-0000-000000000000	100000	HX2W+CHQ London, United Kingdom	40.7128, -74.0060	New York	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,https://example.com/image1.jpg	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-01-01 00:00:00+00	2023-10-27 13:44:38.69121+00	2023-11-01 17:51:26.903279+00	\N	\N
17	9d3468bc-b814-4913-884d-69c13e19f768	Active	00000000-0000-0000-0000-000000000000	20000	address 1	\N	\N		\N	House	1700997025020471429.png,1700997025029965323.webp	12	10	4	best property out there	C	f	Gym	Furnished	2023-12-26 11:04:34+00	2023-11-26 11:10:25.036215+00	2023-11-26 11:10:25.036215+00	38.8951	-77.0364
18	3a92df77-ed96-4ecb-a255-d9e3d3ead328	Active	00000000-0000-0000-0000-000000000000	2000	address 1	\N	\N		\N	Apartment	1700997280794817354.png,1700997280806982032.png	12	2	5	best property out there	C	f	Appliances	Furnished	2023-11-26 11:13:41.613+00	2023-11-26 11:14:40.810585+00	2023-11-26 11:14:40.810585+00	38.8951	-77.0364
19	9178b807-b1ea-4ca4-9652-06f4661e7b7d	Active	00000000-0000-0000-0000-000000000000	25000	address 2003	\N	\N		\N	House	1701804478141877219.png,1701804478149893364.png	12	6	5	the property to be a detail	C	f	Parking	Furnished	2024-01-01 13:15:35+00	2023-12-01 13:17:53.034513+00	2023-12-05 19:27:58.161239+00	38.8951	-77.0364
3	de7eea7d-bd22-4b81-b718-e23d2569934d	Cancelled	00000000-0000-0000-0000-000000000000	100000	HX2W+CHQ London, United Kingdom	40.7128, -74.0060	location	10001	Main St	House	https://example.com/image1.jpg,https://example.com/image2.jpg	2	4	3	Spacious house with a backyard	C	f	Parking	Unfurnished	0001-01-01 00:00:00+00	2023-10-27 13:50:22.303629+00	2023-11-12 19:35:57.068313+00	\N	\N
2	6d03e829-8160-4a0b-a257-994131c309e6	Active	00000000-0000-0000-0000-000000000000	100000	HX2W+CHQ London, United Kingdom	40.7128, -74.0060	location	10001	Main St	House	http://localhost:8001/files/Image_created_with_a_mobile_phone.png,http://localhost:8001/files/istockphoto-1435998461-170667a.webp,http://localhost:8001/files/pexels-photo-106399.jpeg,http://localhost:8001/files/photo-1560518883-ce09059eeffa.avif	2	4	3	Spacious house with a backyard	C	f	Parking,Garden	Unfurnished	0001-01-01 00:00:00+00	2023-10-27 13:48:44.643993+00	2023-10-30 20:16:59.535973+00	\N	\N
\.


--
-- Data for Name: real_estates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.real_estates (id, real_estate_id, email, address, phone_number, why_joined, preferred_time, status, created_at, updated_at) FROM stdin;
3	48bfc073-ad65-4388-99fc-a51fe4a1265f	nahomasnake12@gmail.com		+441234567890	0	2023-10-25 10:00:00+00	ACTIVE	2023-10-24 17:36:14.031796+00	2023-10-24 17:36:14.031796+00
4	a0ee526a-a985-4c41-8765-516eaf318586	nahomasnake12@gmail.com		+441234567890	0	2023-10-25 10:00:00+00	ACTIVE	2023-10-24 17:38:58.25718+00	2023-10-24 17:38:58.25718+00
5	7c59298f-0579-439b-9202-50291587b167	nahomasnake12@gmail.com		+441234567890	1	2023-10-25 10:00:00+00	ACTIVE	2023-10-24 17:45:36.145671+00	2023-10-24 17:45:36.145671+00
6	dd31f4b3-6c63-4ddc-8afa-b09d42a4577b	nahomasnake12@gmail.com		+441234567890	1	2023-10-25 10:00:00+00	ACTIVE	2023-10-24 18:19:51.73142+00	2023-10-24 18:19:51.73142+00
7	8b18b305-f5d0-4028-a1e8-454e3688b904	nahomasnake12@gmail.com		+441234567823	1	2023-10-24 21:00:00+00	ACTIVE	2023-10-24 18:32:49.83585+00	2023-10-24 18:32:49.83585+00
8	ccd4f817-e2a4-484c-b9ba-3f5b7a2c8bc8	nahomasnake12@gmail.com		+441234567823	1	2023-10-24 21:00:00+00	ACTIVE	2023-10-24 18:34:23.32603+00	2023-10-24 18:34:23.32603+00
9	5652ca3a-611a-4052-bdac-6d80678a184b	nahomasnake12@gmail.com		+441234567890	2	2023-10-26 09:00:00+00	ACTIVE	2023-10-24 18:36:12.348672+00	2023-10-24 18:36:12.348672+00
11	ddb829fa-7c86-4318-af61-df7b2c93c4bf	nahomasnake12@gmail.com	171 Union Street, Middlesbrough, North Yorkshire	+441234567123	3	2023-10-27 09:05:00+00	ACTIVE	2023-10-24 19:13:09.302203+00	2023-10-24 19:13:09.302203+00
12	2669da12-3cd0-4ba6-8433-4d0be8a2cafb	nahomasnake12@gmail.com	173 Union Street, Middlesbrough, North Yorkshire	+441234567890	2	2023-10-25 21:00:00+00	ACTIVE	2023-10-24 22:04:53.092664+00	2023-10-24 22:04:53.092664+00
1	580b4768-9cab-471a-8148-256207e4700b	nahomasnake12@gmail.com	randome user address	1441243513	1	2023-10-25 13:52:21.155+00	ACTIVE	2023-10-24 13:54:07.613691+00	2023-10-26 22:20:36.420816+00
2	a85f1b46-a756-4248-b340-a45108a59ffd	nahomasnake12@gmail.com	Address 12, LimbCity	+444912343213	0	2023-11-12 15:14:48.42+00	ACTIVE	2023-10-24 15:04:43.867629+00	2023-11-11 15:33:28.755563+00
10	971dede7-6dd4-4378-b1a1-b4b074d9b617	nahomasnake12@gmail.com	173 Union Street, Middlesbrough, North Yorkshire	+441234567890	2	2023-11-20 09:00:00+00	ACTIVE	2023-10-24 18:38:03.78502+00	2023-11-14 21:48:38.962125+00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, user_id, first_name, middle_name, last_name, email, enrollment, phone, password, user_name, gender, status, created_at, updated_at, role, reset_code) FROM stdin;
18	c2dfcf9a-b41c-46e7-bc89-6c9a2704222b	test	d	user	person1@gmail.com	\N	25190314215	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G			ACTIVE	2023-10-18 18:27:08.958443+00	2023-10-18 18:27:08.958443+00	REGULAR_USER_ROLE	0
1	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	2023-10-17 18:00:34.853047+00	2023-10-25 18:55:00.902166+00	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G	850016
19	336fa10e-a0b6-44d0-9dbb-074c025e2f2f	test	d	user	person2@gmail.com	\N	25190314215	$2a$14$8Hk5SJ//rJ.lgMzeJ5S39.iK4w9BXCSBCo1qzjYg1PYeCmB8lTzFK			ACTIVE	2023-10-19 23:07:34.965235+00	2023-10-19 23:07:34.965235+00	REGULAR_USER_ROLE	0
20	4f2a669f-9012-4505-8992-e8f72a6aa5a8	nahom	asnake	desalegn	nahomsara@gmail.com	\N	+251901051649	$2a$14$4.fa/rfBGGzbSHSG/6XcfOoZSi/oCzW72Ra5ejOflDVJD9zfqdJdG			ACTIVE	2023-10-19 23:55:02.589625+00	2023-10-19 23:55:02.589625+00	REGULAR_USER_ROLE	0
69	2fffde8a-f0e3-483c-91b8-854d141e0f9b	nahom	asnake	desalegn	nahomasnake12@gmail.com	\N	+251901051649	$2a$14$.Z5PZ6RK1WyFK5RMiAWca.GBMz/Hy1RK57AGvaGR3OhR9QaRUJG3G			ACTIVE	2023-10-20 12:03:43.452706+00	2023-10-20 12:03:43.452706+00	ADMIN_ROLE	0
\.


--
-- Name: properties_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.properties_id_seq', 19, true);


--
-- Name: real_estates_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.real_estates_id_seq', 12, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 78, true);


--
-- Name: properties properties_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_pkey PRIMARY KEY (id);


--
-- Name: real_estates real_estates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.real_estates
    ADD CONSTRAINT real_estates_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_email_phone; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_email_phone ON public.users USING btree (email, phone);


--
-- Name: unique_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX unique_email ON public.users USING btree (email);


--
-- PostgreSQL database dump complete
--

