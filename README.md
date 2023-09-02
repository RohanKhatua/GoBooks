# Robust Online Book Store API
## Introduction

This project was made for the Summer Internship Recruitment Process of Balkan ID. We expose a RESTful API which manages an online book store capable of handling user authentication, authorization and access controlled interaction with book store resources.

[Github](https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-RohanKhatua)

🔗The API is hosted [here](https://books-api.rohankhatua.dev/api)

**The Complete API documentation can be found in the form of Postman Collection [here](https://documenter.getpostman.com/view/25992245/2s9Y5ZvMbf).**

## System Design 

![System Design Image](https://raw.githubusercontent.com/RohanKhatua/fiber-jwt/main/sys_design.png?token=GHSAT0AAAAAACE3PL3XLRGCIKAOZLMJADOIZHTG7CA)

## Tech Stack

1. **Go**: Go is a statically-typed, compiled language used for building efficient and scalable backend APIs.
2. **Fiber**: Fiber is a web framework for Go that offers fast and efficient routing and middleware capabilities, making it an excellent choice for building high-performance backend APIs.
3. **Gorm**: Gorm is an Object-Relational Mapping (ORM) library for Go, which simplifies database interactions and allows you to work with your database models in a more Go-like way, enhancing database operations for your backend API.
4. **JWT (JSON Web Token)**: JWT is a compact, self-contained way of securely transmitting information between parties, commonly used for API authentication and authorization, ensuring secure access to your backend API.
5. **PostgreSQL**: PostgreSQL is a powerful open-source relational database system, ideal for storing and managing data for your API.
6. **Docker**: Docker is a containerization platform that simplifies packaging and deploying your API in isolated environments.
7. **AWS S3**: AWS S3 (Amazon Simple Storage Service) is a scalable object storage service that can be used to securely store and serve media assets for your API.
8. **AWS EC2**: AWS EC2 (Amazon Elastic Compute Cloud) provides scalable virtual servers that you can use to host and run your backend API, ensuring reliable and flexible infrastructure.
9. **Nginx**: Nginx is a web server and reverse proxy that can be used to improve the security and performance of your backend API, serving as a powerful gateway for incoming requests.

## Usage

1. Visit the API URL and make requests using the provided API documentation. 
2. Applications like **Postman** can be used to easily send requests and receive responses from the API.
3. `cURL` requests can also be sent to the API. Install `cURL` from [here](https://everything.curl.dev/get).
## Features

### Authentication and Authorization

1. A user can be one of two types - `USER` or `ADMIN`. The `role` of the user is determined when the user signs up.
2. Passwords are hashed using the `SHA 256` algorithm before being stored in the database. When the user enters their password while logging in, the attempted password is hashed using the same algorithm and compared to the stored hash.
3. A token is generated using JWT which signs all of the user's details except their password using a `super_secret` key. This token expires after 24 hours.
4. Requests sent to any route of the API must contain an `Authorization` header containing this token in the form of `Bearer <token>`. This protection is guaranteed by a `JWT Middleware`.
5. The middleware verifies the token and stores the user's details in the context of the application so that they can be used by succeeding routes.
6. Routes which should only be accessed by Administrators check whether the user has the correct role, thus ensuring **RBAC***.
### Account Activation and Deactivation

1. User's can mark their account for deletion or deactivation and can re-activate the account at a later stage. 
2. When deactivated the user cannot perform any actions. 
3. This control is guaranteed by the `Activation Middleware` which sits after the `JWT Middleware`
### Purchases

1. Users can purchase a specified quantity of a certain book provided it exists and the quantity is available.
2. Users can retrieve a list of all purchases made
### Cart Management 

1. The API allows users to add and remove books from the cart from amongst books present in the bookstore.
### Search

1. The user can search for books present in the bookstore by the title or author of the book.
### Reviews

1. Users can leave reviews on books they have purchased.
2. Reviews can be edited and deleted.
3. The average rating of a book can also be accessed.

## Deployment

1. The API is hosted on an AWS EC2 instance.
2. The application is completely Dockerized
3. An Ngnix reverse proxy sits between the user and web server and forwards requests to it using a defined https configuration.