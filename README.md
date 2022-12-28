# Go Bookstore

This is a sample application that demonstrates how to build a Go application using the Clean Architecture.

## Table of Contents

- [Getting Started](#getting-started)
  - [Installing](#installing)
  - [Setup](#Setup)
  - [ER Diagram](#er-diagram)
  - [Router](#installing)
    - [User Router](#user-router)
    - [Book Router](#book-router)
- [Project Structure](#project-structure)

## Getting Started

### ER Diagram

  <img src="./assets/ERD.png">

### Installing

```bash
git clone https://github.com/nutthanonn/go-bookstore.git
cd go-bookstore
go get
```

### Setup

_Create .env file in root project_

```
DB_USERNAME=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
```

# Router

## User Router

|    API Path    | Method |       What it does        |
| :------------: | :----: | :-----------------------: |
|   /api/users   |  GET   | Fetches the list of users |
|   /api/users   |  POST  |        Create user        |
| /api/users/:id | DELETE |        Delete user        |

## Book Router

|    API Path    | Method |       What it does        |
| :------------: | :----: | :-----------------------: |
|   /api/books   |  GET   | Fetches the list of books |
| /api/books/:id |  GET   |     Fetch book by id      |
|   /api/books   |  POST  |        Create book        |
| /api/books/:id |  PUT   |        Update book        |
| /api/books/:id | DELETE |        Delete book        |

## Employee Router

|     API Path      | Method |         What it does          |
| :---------------: | :----: | :---------------------------: |
|   /api/employee   |  GET   | Fetches the list of employees |
| /api/employee/:id |  GET   |     Fetch employee by id      |
|   /api/employee   |  POST  |        Create employee        |
| /api/employee/:id |  PUT   |        Update employee        |
| /api/employee/:id | DELETE |        Delete employee        |

## Project Structure

### /pkg

```bash
.
|____interface
| |____repository
| | |____employee_repository.go
| | |____user_repository.go
| | |____book_repository.go
| |____presenter
| | |____book_presenter.go
| | |____user_presenter.go
| | |____employee_presenter.go
| |____controller
| | |____book_controller.go
| | |____user_controller.go
| | |____employee_controller.go
| | |____app_controller.go
|____registry
| |____employee_registry.go
| |____book_registry.go
| |____registry.go
| |____user_registry.go
|____usecase
| |____repository
| | |____emoployee_repository.go
| | |____user_repository.go
| | |____book_repository.go
| |____presenter
| | |____book_presenter.go
| | |____user_presenter.go
| | |____employee_presenter.go
| |____interactor
| | |____employee_interactor.go
| | |____user_ interactor.go
| | |____book_interactor.go
|____entities
| |____order_details_entities.go
| |____user_entities.go
| |____sale_entities.go
| |____employee_entities.go
| |____book_entities.go
| |____customer_entities.go
| |____order_entities.go
| |____inventory_entities.go
```

### /api

```bash
.
|____routers
| |____employee_handlers.go
| |____book_routers.go
| |____user_routers.go
|____test
| |____employee_router_test.go
|____app.go
|____infrastructure
| |____datastore
| | |____db.go
|____gomon_spawn
|____handlers
| |____employee_handler
| | |____getEmployee.go
| | |____handlers.go
| | |____getEmployeeById.go
| | |____createEmployee.go
| | |____deleteEmployee.go
| | |____updateEmployee.go
| |____book_handler
| | |____getBookById.go
| | |____getBook.go
| | |____handlers.go
| | |____createBook.go
| | |____updateBook.go
| | |____deleteBook.go
| |____user_handler.go
```
