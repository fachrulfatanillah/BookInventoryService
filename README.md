# BookInventoryService API Documentation

Book-Inventory-Service adalah RESTful API untuk mengelola buku, kategori, pengguna dengan autentikasi JWT dan middleware.

---

## Authentication
Semua endpoint categories dan books membutuhkan **JWT token** di header:

## **1. Users**

### **1.1 Register User**
- **Endpoint:** `POST /api/users/register`
- **Body (form-data):**
  - `username` (string, required)
  - `password` (string, required)
  - `created_by` (string, required)
- **Response:**
```json
{
    "message": "User registered successfully",
    "data": {
        "username": "user1"
    }
}
```

### **1.2 Login User**
- **Endpoint:** POST `POST /api/users/register`
- **Body (form-data):**

  - `username` (string, required)
  - `password` (string, required)
- **Response:**
```json
{
  "message": "Login successful",
  "token": "JWT_TOKEN_HERE",
  "username": "user1"
}
```


## **2. Categories**

### **2.1 Create Category**
- **Endpoint:** POST `POST /categories`
- **Body (form-data):**

  - name (string, required)
- **Response:**
```json
{
  "message": "Category created successfully",
  "data": {
    "id": 1,
    "name": "Fiction",
    "created_at": "2025-12-08T10:00:00Z",
    "created_by": "admin",
    "modified_at": "2025-12-08T10:00:00Z",
    "modified_by": "admin"
  }
}
```

### **2.2 Get All Categories**
- **Endpoint:** GET `GET /api/categories`
- **Response:**
```json
{
  "message": "Categories fetched successfully",
  "data": [
    {
      "id": 1,
      "name": "Fiction",
      "created_at": "2025-12-08T10:00:00Z",
      "created_by": "admin",
      "modified_at": "2025-12-08T10:00:00Z",
      "modified_by": "admin"
    }
  ]
}
```

### **2.3 Get Category by ID**
- **Endpoint:** GET `GET /api/categories/:id`
- **Response:**
```json
{
  "message": "Success",
  "category": {
    "id": 1,
    "name": "Fiction",
    "created_at": "2025-12-08T10:00:00Z",
    "created_by": "admin",
    "modified_at": "2025-12-08T10:00:00Z",
    "modified_by": "admin"
  }
}
```

### **2.4 Update Category**
- **Endpoint:** PUT `PUT /api/categories/:id`
- **Body (form-data):**

  - `name` (string, required)

- **Response:**
```json
{
  "message": "Category updated successfully",
  "data": {
    "id": 1,
    "name": "Science Fiction",
    "created_at": "2025-12-08T10:00:00Z",
    "created_by": "admin",
    "modified_at": "2025-12-08T12:00:00Z",
    "modified_by": "admin"
  }
}
```

### **2.5 Delete Category (beserta semua buku)**
- **Endpoint:** DELETE `DELETE /api/categories/:id`

- **Response:**
```json
{
  "message": "Category and all books under it deleted successfully"
}
```

### **2.6 Get Books by Category**
- **Endpoint:** GET `GET /api/categories/:id/books`

- **Response:**
```json
{
  "message": "Books fetched successfully",
  "category": "Fiction",
  "data": [
    {
      "id": 1,
      "title": "The Go Programming Language",
      "description": "Belajar bahasa Go",
      "image_url": "https://example.com/image.jpg",
      "release_year": 2015,
      "price": 100000,
      "total_page": 350,
      "thickness": "tebal",
      "category_id": 1,
      "created_at": "2025-12-08T10:00:00Z",
      "created_by": "admin",
      "modified_at": "2025-12-08T10:00:00Z",
      "modified_by": "admin"
    }
  ]
}
```

## **3. Books**

### **3.1 Create Book**
- **Endpoint:** POST `POST /api/books`
- **Body (form-data):**

  - `title` (string, required)
  - `description` (string, required)
  - `image_url` (string, optional)
  - `release_year` (int, required, 1980-2024)
  - `price` (int, optional)
  - `total_page` (int, required)
  - `category_id` (int, required)

- **Response:**
```json
{
  "message": "Book created successfully",
  "data": {
    "id": 1,
    "title": "The Go Programming Language",
    "description": "Belajar bahasa Go",
    "image_url": "https://example.com/image.jpg",
    "release_year": 2015,
    "price": 100000,
    "total_page": 350,
    "thickness": "tebal",
    "category_id": 1,
    "created_at": "2025-12-08T10:00:00Z",
    "created_by": "admin",
    "modified_at": "2025-12-08T10:00:00Z",
    "modified_by": "admin"
  }
}
```

### **3.2 Create Book**
- **Endpoint:** POST `POST /api/books`

- **Response:**
```json
{
  "message": "Books retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "The Go Programming Language",
      "description": "Belajar bahasa Go",
      "image_url": "https://example.com/image.jpg",
      "release_year": 2015,
      "price": 100000,
      "total_page": 350,
      "thickness": "tebal",
      "category_id": 1,
      "created_at": "2025-12-08T10:00:00Z",
      "created_by": "admin",
      "modified_at": "2025-12-08T10:00:00Z",
      "modified_by": "admin"
    }
  ]
}
```

### **3.3 Get Book by ID**
- **Endpoint:** GET `GET /api/books/:id`

- **Response:**
```json
{
  "message": "Book retrieved successfully",
  "data": {
    "id": 1,
    "title": "The Go Programming Language",
    "description": "Belajar bahasa Go",
    "image_url": "https://example.com/image.jpg",
    "release_year": 2015,
    "price": 100000,
    "total_page": 350,
    "thickness": "tebal",
    "category_id": 1,
    "created_at": "2025-12-08T10:00:00Z",
    "created_by": "admin",
    "modified_at": "2025-12-08T10:00:00Z",
    "modified_by": "admin"
  }
}
```

### **3.4 Edit Book**
- **Endpoint:** PUT `PUT /api/books/:id`
- **Body (form-data):**

  - Bisa update salah satu atau semua field: `title`, `description`, `image_url`, `release_year`, `price`, `total_page`, `category_id`

- **Response:**
```json
{
  "message": "Book updated successfully",
  "data": {
    "id": 1,
    "title": "The Go Programming Language Updated",
    "description": "Belajar bahasa Go lebih lanjut",
    "image_url": "https://example.com/image.jpg",
    "release_year": 2016,
    "price": 120000,
    "total_page": 360,
    "thickness": "tebal",
    "category_id": 1,
    "created_at": "2025-12-08T10:00:00Z",
    "created_by": "admin",
    "modified_at": "2025-12-08T12:00:00Z",
    "modified_by": "admin"
  }
}
```

### **3.5 Delete Book**
- **Endpoint:** DELETE `DELETE /api/books/:id`

- **Response:**
```json
{
  "message": "Book deleted successfully"
}
```  