# Golang Library REST API

Library REST API made with Go v1.23.2

### Users API endpoints:

- `GET /users/list`
- `POST /users/add => Example payload { "username": (value string), "email": (value string) }`

### Borrowed books API endpoints:

- `GET /borrowed/{userId}`
- `GET /returned/{userId}`
- `POST /borrowed/add => Example payload { "book_id": (value int), "user_id": (value int), "borrow_date": (value string[YYYY-MM-DD hh:mm:ss]) }`
- `PUT /returned/update => Example payload { "id": (value int), "return_date": (value string[YYYY-MM-DD hh:mm:ss]) }`

### Books API endpoints:

- `GET /books/list`

## Database

Create PostgreSQL database

`CREATE DATABASE library ENCODING UTF-8;`

## Project setup

Create .env file in project root folder

`cp .env.example .env`

and update Database credentials

Install [Golang Migrate](https://github.com/golang-migrate/migrate)

Setup Go dependencies

`go get .`

`go mod tidy`

Run Go project

`go run main.go`

After executing the command, Migrations will run, which will setup your database tables and mock data (code in main.go, main() method)
