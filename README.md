## REST API Written In Go

> This simple REST API was built to better understand the development of "CRUD" applications

## Main learning points:

* how to work with Postgresql
* writing unit tests for the REST API endpoints
* how to structure a Go project
* how to access environment variables
* how to set-up and use Docker 

## Dependencies:

* github.com/gorilla/mux v1.8.0
*	github.com/joho/godotenv v1.4.0
*	github.com/lib/pq v1.10.4

## How To Satrt The Application

```bash 
create database - psql \i /<your>/<path>/restapi/schema.sql
git clone https://github.com/ElmanZ/rest-api.git
cd go/src/github.com/<github username>/<project name>
create and edit .env file to specify your env variables
docker-compose up --build
```

## Examples

POST(creates new data) Request:

```bash
(unix shell) - curl --data '{"username": "Mindy"}' http://localhost:9000/user/add
(windows prompt) - curl --data "{\"username\": \"Mindy\"}" http://localhost:9000/user/add
```
Response: {"id":1,"username":"Mindy"}

POST(creates new data) Request:

```bash
(unix shell) - curl --data '{"name": "chat_1", "users": 1}' http://localhost:9000/chats/add
(windows prompt) - curl --data "{\"name\": \"chat_1\", \"users\": 1}" http://localhost:9000/chats/add

```
Response: {"id":1,"name":"chat_1","users":1}

GET(retrieves created data) Request:
```bash
(unix/windows) - curl localhost:9000/chats/get/1
```
Response: {"id":1,"name":"chat_1","users":1}

PUT(updates the data) Request:
```bash
(unix shell) - curl -X PUT --data '{"username": "Poppy"}' http://localhost:9000/user/update/1
(windows prompt) - curl -X PUT --data "{\"username\": \"Poppy\"}' http://localhost:9000/user/update/1

```
Response: {"id":1,"username":"Poppy"}

DELETE(deletes data) Request:
```bash
(unix/windows) - curl -X DELETE localhost:9000/chats/delete/4 
```
Response: "chat was deleted successfully"

## Test
```bash
ez@MBP db % go test -v
=== RUN   TestAddUser
--- PASS: TestAddUser (0.01s)
=== RUN   TestAddChat
--- PASS: TestAddChat (0.00s)
=== RUN   TestUpdateUser
--- PASS: TestUpdateUser (0.00s)
=== RUN   TestGetChat
--- PASS: TestGetChat (0.00s)
=== RUN   TestDeleteChat
--- PASS: TestDeleteChat (0.00s)
PASS
ok      /restapi/db    0.482s
```