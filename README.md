# restfulApiTest

*This is a RESTful JSON API test in golang.*

## Install
go get github.com/jhonsky/restfulApiTest

Use RESTful lib:[![go-json-rest](https://github.com/ant0ine/go-json-rest.git)].

Use postgreSql lib:[![lib-pq](https://github.com/lib/pq.git)].

Use database postgreSql.


The following is database design:

DBName=tantan

first table: tbl_user

create tbl_user sql:
```
CREATE TABLE public.tbl_user
(
  id integer NOT NULL DEFAULT nextval('tbl_user_id_seq'::regclass),
  name character varying,
  CONSTRAINT tbl_user_pkey PRIMARY KEY (id)
)
```
second table: tbl_relation

create tbl_relation sql:
```
CREATE TABLE public.tbl_relation
(
  id integer NOT NULL DEFAULT nextval('tbl_relation_id_seq'::regclass),
  usera integer NOT NULL,
  userb integer NOT NULL,
  state character varying,
  CONSTRAINT tbl_relation_pkey PRIMARY KEY (id)
)
```