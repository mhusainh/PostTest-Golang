# Weekly-Task3: Bookstore RESTful API

![Go](https://img.shields.io/badge/Go-1.20-blue.svg)
![GORM](https://img.shields.io/badge/GORM-v2.0.0-green.svg)
![SQLite](https://img.shields.io/badge/SQLite-3.36.0-blue.svg)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)

## Table of Contents

- [About](#about)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Running the Server](#running-the-server)
  - [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## About

**Weekly-Task3** is a simple RESTful API built with Go (Golang) for managing a bookstore's inventory. The API allows users to perform CRUD (Create, Read, Update, Delete) operations on book records stored in a SQLite database. It leverages GORM v2 for ORM (Object-Relational Mapping) and Gorilla Mux for HTTP request routing.

## Features

- **Create a Book**: Add new books to the inventory.
- **Retrieve Books**: Get a list of all books or a specific book by ID.
- **Update a Book**: Modify details of an existing book.
- **Delete a Book**: Remove a book from the inventory.
- **JSON Responses**: All API responses are in JSON format.
- **Error Handling**: Consistent and informative error responses.

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [GORM v2](https://gorm.io/) - ORM library for Go
- [SQLite](https://www.sqlite.org/index.html) - Lightweight relational database
- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP request router and dispatcher
- [Postman](https://www.postman.com/) - API testing tool (optional)

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: Make sure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).
- **Git**: To clone the repository. Download from [here](https://git-scm.com/downloads).

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/Weekly-Task3.git
   ```

2. **Navigate to the Project Directory**

   ```bash
   cd Weekly-Task3
   ```

3. **Initialize Go Modules**

   Ensure you are in the root directory of the project and run:

   ```bash
   go mod tidy
   ```

   This will download all necessary dependencies as specified in `go.mod`.

## Usage

### Running the Server

1. **Build and Run the Application**

   ```bash
   go run main.go
   ```

   Alternatively, you can build the application and then run the executable:

   ```bash
   go build -o bookstore
   ./bookstore
   ```

2. **Server Output**

   After running the server, you should see:

   ```
   Server running on http://localhost:8080
   ```

### API Endpoints

The API provides the following endpoints for managing books:

#### 1. Create a New Book

- **Endpoint**: `POST /book/`
- **Description**: Adds a new book to the inventory.
- **Request Body**:

  ```json
  {
    "title": "Belajar Golang",
    "author": "Budi",
    "price": 150000
  }
  ```

- **Response**:
  - **Status**: `201 Created`
  - **Body**:

    ```json
    {
      "id": 1,
      "title": "Belajar Golang",
      "author": "Budi",
      "price": 150000
    }
    ```

#### 2. Retrieve All Books

- **Endpoint**: `GET /book/`
- **Description**: Retrieves a list of all books.
- **Response**:
  - **Status**: `200 OK`
  - **Body**:

    ```json
    [
      {
        "id": 1,
        "title": "Belajar Golang",
        "author": "Budi",
        "price": 150000
      }
    ]
    ```

#### 3. Retrieve a Book by ID

- **Endpoint**: `GET /book/{bookid}`
- **Description**: Retrieves details of a specific book by its ID.
- **Parameters**:
  - `bookid` (path) - ID of the book to retrieve.
- **Response**:
  - **Status**: `200 OK`
  - **Body**:

    ```json
    {
      "id": 1,
      "title": "Belajar Golang",
      "author": "Budi",
      "price": 150000
    }
    ```

#### 4. Update a Book

- **Endpoint**: `PUT /book/{bookid}`
- **Description**: Updates details of an existing book.
- **Parameters**:
  - `bookid` (path) - ID of the book to update.
- **Request Body**:

  ```json
  {
    "title": "Belajar Golang Lanjutan",
    "author": "Budi",
    "price": 175000
  }
  ```

- **Response**:
  - **Status**: `200 OK`
  - **Body**:

    ```json
    {
      "id": 1,
      "title": "Belajar Golang Lanjutan",
      "author": "Budi",
      "price": 175000
    }
    ```

#### 5. Delete a Book

- **Endpoint**: `DELETE /book/{bookid}`
- **Description**: Deletes a book from the inventory.
- **Parameters**:
  - `bookid` (path) - ID of the book to delete.
- **Response**:
  - **Status**: `204 No Content`

### Testing the API

You can use tools like [Postman](https://www.postman.com/) or `curl` to test the API endpoints.

**Example using `curl`:**

- **Create a New Book**

  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"title":"Belajar Go", "author":"John Doe", "price":99.99}' http://localhost:8080/book/
  ```

- **Retrieve All Books**

  ```bash
  curl -X GET http://localhost:8080/book/
  ```

- **Retrieve a Book by ID**

  ```bash
  curl -X GET http://localhost:8080/book/1
  ```

- **Update a Book**

  ```bash
  curl -X PUT -H "Content-Type: application/json" -d '{"title":"Belajar Go Lang", "author":"John Doe", "price":89.99}' http://localhost:8080/book/1
  ```

- **Delete a Book**

  ```bash
  curl -X DELETE http://localhost:8080/book/1
  ```

## Project Structure

```
Weekly-Task3/
├── main.go
├── go.mod
├── go.sum
├── README.md
└── pkg/
    ├── config/
    │   └── app.go
    ├── controllers/
    │   └── controllers.go
    ├── models/
    │   └── book.go
    ├── routes/
    │   └── routes.go
    └── utils/
        └── utils.go
```

- **main.go**: Entry point of the application. Initializes the database connection, sets up routes, and starts the server.
- **pkg/config/app.go**: Handles database connection and migration using GORM.
- **pkg/models/book.go**: Defines the `Book` model representing the book entity.
- **pkg/controllers/controllers.go**: Contains handler functions for API endpoints.
- **pkg/routes/routes.go**: Sets up API routes and maps them to controller functions.
- **pkg/utils/utils.go**: Provides utility functions for error handling and JSON responses.

## Contributing

Contributions are welcome! Please follow these steps:

1. **Fork the Repository**

   Click the "Fork" button at the top right of this page to create your own fork.

2. **Clone Your Fork**

   ```bash
   git clone https://github.com/yourusername/Weekly-Task3.git
   ```

3. **Create a New Branch**

   ```bash
   git checkout -b feature/YourFeatureName
   ```

4. **Make Changes and Commit**

   ```bash
   git add .
   git commit -m "Add Your Feature"
   ```

5. **Push to Your Fork**

   ```bash
   git push origin feature/YourFeatureName
   ```

6. **Create a Pull Request**

   Go to the original repository and create a pull request from your forked repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

If you have any questions or feedback, feel free to reach out:

- **Name**: Mahmudi Husain Hasbullah
- **Email**: husainhasbulah@gmail.com
- **GitHub**: [mhusainh](https://github.com/mhusainh)

---

Thank you for checking out **Weekly-Task3**! We hope this project serves as a useful reference for building RESTful APIs with Go. Happy coding!