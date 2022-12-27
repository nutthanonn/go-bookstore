# Go Bookstore

This is a sample application that demonstrates how to build a Go application using the Clean Architecture.

## Table of Contents

- [Getting Started](#getting-started)
  - [Installing](#installing)
  - [ER Diagram](#er-diagram)
  - [Router](#installing)
    - [User Router](#user-router)
    - [Book Router](#book-router)
- [Project Structure](#project-structure)

## Getting Started

<!-- ### Prerequisites -->

### ER Diagram

  <img src="./assets/ERD.png">

### Installing

```bash
git clone https://github.com/nutthanonn/go-bookstore.git
cd go-bookstore
go get
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

## Project Structure

### /pkg

```bash
.
|____interface
| |____repository
| | |____user_repository.go
| | |____book_repository.go
| |____presenter
| | |____book_presenter.go
| | |____user_presenter.go
| |____controller
| | |____book_controller.go
| | |____user_controller.go
| | |____app_controller.go
|____registry
| |____book_registry.go
| |____registry.go
| |____user_registry.go
|____usecase
| |____repository
| | |____user_repository.go
| | |____book_repository.go
| |____presenter
| | |____book_presenter.go
| | |____user_presenter.go
| |____interactor
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
| |____book_routers.go
| |____user_routers.go
|____app.go
|____infrastructure
| |____datastore
| | |____db.go
|____gomon_spawn
|____handlers
| |____book_handler
| | |____getBookById.go
| | |____getBook.go
| | |____handlers.go
| | |____createBook.go
| | |____updateBook.go
| | |____deleteBook.go
| |____user_handler.go
```
