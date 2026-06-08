# Book Rental
This project is a REST API for a book rental system developed using Golang and the Echo Framework. It uses GORM as the ORM, PostgreSQL as the database, JWT for authentication and authorization, and Swagger for API documentation.

# Feature
- User Register
- User Login
- Get User Profile
- JWT Authentication
- Swagger Documentation

# Tech Stack
- Golang
- Echo Framework
- PostgreSQL
- GORM
- JWT
- Swagger

# Instalation
- Clone repository:
  - git clone <repository-url>
  - cd book-rental-api

- Copy environment file:
  - cp .env.example .env
  - Install dependencies:
  - go mod tidy

- Run application:
  - go run main.go

# Environment Variables
- Create .env file based on .env.example
    - DB_HOST=localhost
    - DB_PORT=5432
    - DB_USER=postgres
    - DB_PASSWORD=password
    - DB_NAME=book_rental

    - JWT_SECRET=secret
    - PORT=8080

# API Documentation
- Swagger: http://localhost:8080/swagger/index.html
- Deployed url: book-rental-api-production-8c4a.up.railway.app
- The Web API can be accessed at: https://book-rental-api-production-8c4a.up.railway.app/swagger/index.html

# End Point
Users
Method	Endpoint	Description
POST	/users/register	Register user
POST	/users/login	Login user
GET	/users/profile	Get user profile

#### Author
Raden Tio Genta Komara (Komara)
