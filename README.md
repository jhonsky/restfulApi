# restfulApiTest

*This is a RESTful JSON API test in golang.*

## Install
go get github.com/jhonsky/restfulApiTest

Use RESTful lib:[![go-json-rest](https://github.com/ant0ine/go-json-rest.git)].

Use postgreSql lib:[![lib-pq](https://github.com/lib/pq.git)].

Use database postgreSql.


The following is database design:
```
host := "localhost"
port := 5432
dbID := "tantan"
user := "postgres"
pwd := "1111"
```
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

The following is examples.
```
GET/users 
List all users 
Example: 
$curl -XGET "http://localhost:9090/users" 
[
    {
        "id": "21341231231",
        "name": "Bob",
        "type": "user"
    },
    {
        "id": "31231242322",
        "name": "Samantha",
        "type": "user"
    }
] 
```
```
POST/users 
Create a user 
allowed fields: 
name = string 
Example: 
$curl -XPOST -d '{"name":"Alice"}' "http://localhost:9090/users" 
{ 
	"id": "11231244213", 
	"name": "Alice" ,
	"type": "user" 
}
```
```
GET/users/:user_id/relationships 
List a users all relationships 
Example: 
$curl -XGET "http://localhost:9090/users/11231244213/relationships" 
[
    {
        "user_id": "222333444",
        "state": "liked",
        "type": "relationship"
    },
    {
        "user_id": "333222444",
        "state": "matched",
        "type": "relationship"
    },
    {
        "user_id": "444333222",
        "state": "disliked",
        "type": "relationship"
    }
]
```
```
PUT/users/:user_id/relationships/:other_user_id 
Create/update relationship state to another user. 
allowed fields: 
state = "liked"|"disliked" 
If two users have "liked" each other, then the state of the relationship is "matched" 
Example: 
$curl -XPUT -d '{"state":"liked"}' 
"http://localhost:9090/users/11231244213/relationships/21341231231" 
{ 
	"user_id": "21341231231", 
	"state": "liked" ,
	"type": "relationship" 
} 
$curl -XPUT -d '{"state":"liked"}' 
"http://localhost:9090/users/21341231231/relationships/11231244213" 
{ 
	"user_id": "11231244213", 
	"state": "matched" ,
	"type": "relationship" 
} 
$curl -XPUT -d '{"state":"disliked"}' 
"http://localhost:9090/users/21341231231/relationships/11231244213" 
{ 
	"user_id": "11231244213", 
	"state": "disliked" ,
	"type": "relationship" 
}
```