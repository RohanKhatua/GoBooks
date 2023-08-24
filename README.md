
# Fiber JWT Authentication

This is an example project that demonstrates how to implement user registration, login, and JSON Web Token (JWT) authentication using the Fiber web framework and the jwt-go library in Go (Golang).

## Features

- User registration with password hashing
- User login with JWT token generation
- Role-based authentication (USER and ADMIN)
- Middleware for protecting routes using JWT
- Basic error handling and response structures

## Requirements

- Go (Golang) installed
- Basic understanding of RESTful APIs and web development concepts
- A database system (e.g., PostgreSQL, MySQL) for user data storage

## Installation

1. Clone this repository:

```bash
git clone https://github.com/RohanKhatua/fiber-jwt.git
cd fiber-jwt
```

2. Set up your database and configure the database connection in the `database` package.

3. Install the required dependencies:

```bash
go get -u github.com/gofiber/fiber/v2
go get -u github.com/joho/godotenv
go get -u github.com/dgrijalva/jwt-go
```

4. Create a `.env` and set your `super_secret` value.

5. Run the application:

```bash
go run main.go
```

Or,

```bash
air
```

## Usage

- Register a new user: `POST /signup`
- Login: `POST /login`
- Access protected routes: Set the JWT token in the `Authorization` header as `Bearer <token>`.



