# 📝 Todo List Backend (Golang)

A simple REST API built with **Go (Golang)** that performs basic CRUD operations for managing todos.

> Minimal backend project focused only on CRUD functionality.

---

## 🚀 Features

- ➕ Create a Todo
- 📄 Get All Todos
- 🔍 Get Todo by ID
- ✏️ Update a Todo
- ❌ Delete a Todo

---

## 🛠 Tech Stack

- Go (Golang)
- net/http (or Gin / Chi / Gorilla Mux)
- JSON
- In-memory storage (or database)

---

## ⚙️ Setup & Installation

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/yourusername/todo-backend.git
cd todo-backend
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Run the Server

```bash
go run main.go -config config/local.yaml
```

Server runs on:

```
http://localhost:3000
```

---

## 📌 API Endpoints

### ➕ Create Todo

**POST** `/create`

Request Body:

```json
{
  "title": "Learn Go",
  "description": "i love it",
  "completed": false
}
```

---

### 📄 Get All Todos

**GET** `/all`

---

### 🔍 Get Todo by ID

**GET** `/get/{id}`

---

### ✏️ Update Todo

**PUT** `/update/{id}`

Request Body:

```json
{
  "title": "Learn Go properly",
  "description": "i love it",
  "completed": true
}
```

---

### ❌ Delete Todo

**DELETE** `/delete/{id}`

---

## 📦 Example Todo Model

```go
type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Description string `json:"description"`
    Completed bool   `json:"completed"`
}
```

---

## 📝 Notes

- Basic CRUD implementation
- No authentication
- No advanced validation
- Built for learning purposes
