# 📝 Notes App API

**Notes App** is a personal study project focused on building a RESTful API for managing notes. The API allows creating and storing notes using a PostgreSQL database. The main goal is to practice back-end development concepts using Go, database integration, and API best practices.

## ⚙️ Technologies Used

- **Go** – main programming language of the application;
- **PostgreSQL** – relational database used to store notes;
- **GORM** – ORM used to simplify database interactions;
- **Docker** – used to streamline development and deployment.

## 🚧 Current Features

- [x] Note creation and storage in the database;
- [x] Basic integration between Go and PostgreSQL using GORM.

## 🛣️ Roadmap (Planned Features)

- [ ] User authentication and registration with OAuth2;
- [ ] Listing, editing, and deleting notes;
- [ ] Improved validation and error handling;
- [ ] API documentation using Swagger/OpenAPI.
- [ ] Front-end built with React for interacting with the API;
- [ ] Step-by-step guide using `Docker` to run the project locally.

## 📦 How to Run
### Without Docker
Install [Go](https://go.dev/doc/install) and [Postgres](https://www.postgresql.org/download/) if don't already have it.

Download the project's dependencies and run it:
```
go mod tidy  
go run cmd/main.go
```

### With Docker
A step-by-step guide using `Docker` to run the project locally will be added soon.