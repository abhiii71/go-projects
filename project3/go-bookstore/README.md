
---

# 📚 Go Bookstore REST API

This is a simple RESTful API built in **Go** using the **Gorilla Mux** router and **GORM** ORM with a **MySQL** backend. It allows users to **create, read, update, and delete** books from a database.

---


## 🧱 Architecture: MVC Pattern

This project follows the **Model-View-Controller (MVC)** architecture:

### 1. **Model (pkg/models)**
- Contains the data structure (`Book`) and handles database interaction using GORM.

### 2. **Controller (pkg/controllers)**
- Contains the HTTP handler functions (e.g., `CreateBook`, `GetBookById`, `UpdateBook`).
- Parses request data, invokes model functions, and formats JSON responses.

### 3. **View/Router (pkg/routes)**
- Uses Gorilla Mux to route incoming HTTP requests to the appropriate controller function.

### 4. **Utility (pkg/utils)**
- Contains helper functions like `ParseBody` for parsing request payloads.

### 5. **Main (cmd/main/main.go)**
- Application entry point. Initializes DB connection, sets up routes, and starts the server.

---

## 📦 API Endpoints

| Method | Endpoint         | Description              |
|--------|------------------|--------------------------|
| GET    | `/book/`         | Get all books            |
| GET    | `/book/{id}`     | Get a book by ID         |
| POST   | `/book/`         | Create a new book        |
| PUT    | `/book/{id}`     | Update an existing book  |
| DELETE | `/book/{id}`     | Delete a book            |

---

## 📁 Project Structure

```
.
├── cmd
│   └── main
│       └── main.go             # Entry point
├── pkg
│   ├── config
│   │   └── app.go              # DB connection logic
│   ├── controllers
│   │   └── book-controller.go  # CRUD handlers
│   ├── models
│   │   └── book.go             # Book model
│   ├── routes
│   │   └── bookstore-route.go  # Route declarations
│   └── utils
│       └── utils.go            # JSON parsing helpers
├── go.mod
├── go.sum
└── README.md
```

---

## 🔧 Setup Instructions

### ✅ 1. Install Go Modules

```bash
go get github.com/jinzhu/gorm  
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/gorilla/mux
```

---

### ✅ 2. Set Up MySQL Locally

```bash
sudo apt-get update
sudo apt-get install mysql-server
sudo systemctl start mysql
sudo systemctl enable mysql
```

Then, login to MySQL:

```bash
sudo mysql
```

Create a database and user:

```sql
CREATE DATABASE simplerest;
CREATE USER 'root'@'localhost' IDENTIFIED BY 'yourpassword';
GRANT ALL PRIVILEGES ON simplerest.* TO 'root'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```

---

### ✅ 3. Configure the DB Connection

Update your `pkg/config/app.go` file:

```go
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:yourpassword@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
```

---

## 🚀 Run the Project

```bash
go run cmd/main/main.go
```

Expected output:

```
2025/07/21 18:15:43 Server started on port :9090
```

---

## 📨 API Endpoints

### 📘 1. Get All Books

* **GET** `/book/`

```bash
curl -X GET http://localhost:9090/book/
```

---

### 🆕 2. Create a New Book

* **POST** `/book/`

**JSON Body:**

```json
{
  "name": "The Alchemist",
  "author": "Paulo Coelho",
  "publication": "HarperOne"
}
```

```bash
curl -X POST http://localhost:9090/book/ \
  -H "Content-Type: application/json" \
  -d '{"name":"The Alchemist", "author":"Paulo Coelho", "publication":"HarperOne"}'
```

---

### 📗 3. Get Book by ID

* **GET** `/book/{id}`

```bash
curl -X GET http://localhost:9090/book/1
```

---

### ✏️ 4. Update a Book

* **PUT** `/book/{id}`

**JSON Body:**

```json
{
  "name": "The Alchemist - Updated",
  "author": "Paulo Coelho",
  "publication": "HarperCollins"
}
```

```bash
curl -X PUT http://localhost:9090/book/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"The Alchemist - Updated", "author":"Paulo Coelho", "publication":"HarperCollins"}'
```

---

### 🗑️ 5. Delete a Book

* **DELETE** `/book/{id}`

```bash
curl -X DELETE http://localhost:9090/book/1
```

---

## 🧠 Development Strategy

1. Start with **routes** (`bookstore-route.go`)
2. Configure **DB connection** (`app.go`)
3. Add **utils** for parsing JSON (`utils.go`)
4. Initialize entry point (`main.go`)
5. Create **models** (`book.go`)
6. Implement controller logic (`book-controller.go`)
7. Set up and test with MySQL

---

## 🛠️ Built With

* [Go](https://golang.org/)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [GORM](https://gorm.io/)
* [MySQL](https://www.mysql.com/)

---
