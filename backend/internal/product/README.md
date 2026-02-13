# 📦 Product API Documentation

## 📌 Base URL
- http://127.0.0.1:8080

## 📌 Overview

This API provides CRUD operations for managing products stored in a CSV file.

- Framework: Go + Fiber
- Data Source: `datasets/ecommerce/products.csv`
- Storage: In-memory `map[string]Product`
- Testing Tool: **httpie**

---

## 📌 Product Schema

Each product has the following structure:

```json
{
  "product_id": "string (PROD-xxxxxxxx)",
  "name": "string",
  "category_id": "string",
  "price": "float64",
  "description": "string",
  "created_at": "RFC3339 timestamp"
}
```

---

## 🧪 API Manual Testing (httpie)

Base URL: http://127.0.0.1:8080  
Tool Used: httpie

---

# 1️⃣ GET /products

## Test 1: Retrieve All Products

```bash
http GET http://127.0.0.1:8080/products
```

```json
HTTP/1.1 200 OK
Content-Length: 1650
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:05:51 GMT
Server: backend

[
    {
        "category_id": "CAT-612f3042",
        "created_at": "2026-02-13T16:35:46.271752589+05:30",
        "description": "Out news late nice song spring. Expert agency record book because scientist themselves quality. Consider floor scientist sit. East police game first.",
        "name": "East Mini",
        "price": 850.82,
        "product_id": "PROD-6fceba54"
    },....]
```

---

## Test 2: Empty Product List

```bash
http GET http://127.0.0.1:8080/products
```

```json
HTTP/1.1 200 OK
Content-Length: 2
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:07:29 GMT
Server: backend

[]
```

---

# 2️⃣ GET /products/:id

## Test 1: Valid ID

```bash
http GET http://127.0.0.1:8080/products/PROD-31dfb83c
```

```json
HTTP/1.1 200 OK
Content-Length: 315
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:08:38 GMT
Server: backend

{
    "category_id": "CAT-88044df4",
    "created_at": "2026-02-13T16:37:24.514536363+05:30",
    "description": "Think concern you data any. Attorney close yes. Free various military idea even reality. Study born put enjoy budget move relate. Remain though single.",
    "name": "Always Ultra",
    "price": 268.43,
    "product_id": "PROD-31dfb83c"
}
```

---

## Test 2: Non-Existent ID

```bash
http GET http://127.0.0.1:8080/products/PROD-00000000
```

```json
HTTP/1.1 404 Not Found
HTTP/1.1 404 Not Found
Content-Length: 29
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:09:43 GMT
Server: backend

{
    "error": "Product not found"
}
```

---

## Test 3: Invalid ID Format

```bash
http GET http://127.0.0.1:8080/products/INVALID-ID
```

```json
HTTP/1.1 404 Not Found
Content-Length: 29
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:10:46 GMT
Server: backend

{
    "error": "Product not found"
}
```

---

# 3️⃣ POST /products

## Test 1: Valid Product Creation

```bash
http POST http://127.0.0.1:8080/products \
name="Phone" \
category_id="CAT-2" \
price:=20000 \
description="Android smartphone"
```

```json
HTTP/1.1 201 Created
Content-Length: 167
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:11:12 GMT
Server: backend

{
    "category_id": "CAT-2",
    "created_at": "2026-02-13T16:41:13.633346037+05:30",
    "description": "Android smartphone",
    "name": "Phone",
    "price": 20000,
    "product_id": "PROD-83497471"
}
```

---

## Test 2: Missing Fields

```bash
http POST http://127.0.0.1:8080/products name="Phone"
```

```json
HTTP/1.1 206 Partial Content
Content-Length: 71
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:13:56 GMT
Server: backend

{
    "error": "all Product fields required to be filled for Product creation"
}
```

---

## Test 3: Malformed JSON

```bash
echo '{bad json}' | http POST http://127.0.0.1:8080/products
```

```json
HTTP/1.1 400 Bad Request
Content-Length: 24
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:14:52 GMT
Server: backend

{
    "error": "Invalid JSON"
}
```

---

# 4️⃣ PUT /products/:id

## Test 1: Successful Update

```bash
http PUT http://127.0.0.1:8080/products/PROD-31dfb83c \
name="Updated Phone" \
category_id="CAT-2" \
price:=22000 \
description="Updated description"
```

```json
HTTP/1.1 200 OK
Content-Length: 175
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:15:49 GMT
Server: backend

{
    "category_id": "CAT-2",
    "created_at": "2026-02-13T16:44:22.97480591+05:30",
    "description": "Updated description",
    "name": "Updated Phone",
    "price": 22000,
    "product_id": "PROD-31dfb83c"
}
```

---

## Test 2: Non-Existent ID

```bash
http PUT http://127.0.0.1:8080/products/PROD-00000000 \
name="Test" \
category_id="CAT-1" \
price:=1000 \
description="Test"
```

```json
HTTP/1.1 404 Not Found
Content-Length: 36
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:16:12 GMT
Server: backend

{
    "error": "Product not found"
}
```

---

## Test 3: Incomplete Data

```bash
http PUT http://127.0.0.1:8080/products/PROD-87654321 name="Incomplete"
```

```json
HTTP/1.1 404 Not Found
Content-Length: 36
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:16:41 GMT
Server: backend

{
    "error": "Product not found"
}
```

---

# 5️⃣ DELETE /products/:id

## Test 1: Successful Delete

```bash
http DELETE http://127.0.0.1:8080/products/PROD-31dfb83c 
```

```json
HTTP/1.1 200 OK
Content-Length: 21
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:18:36 GMT
Server: backend

{
    "message": "Deleted"
}
```

---

## Test 2: Delete Already Deleted Product

```bash
http DELETE http://127.0.0.1:8080/products/PROD-31dfb83c 
```

```json
HTTP/1.1 404 Not Found
Content-Length: 43
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:19:13 GMT
Server: backend

{
    "error": "Product to be deleted not found"
}
```

---

## Test 3: Delete Non-Existent ID

```bash
http DELETE http://127.0.0.1:8080/products/PROD-00000000
```

```json
HTTP/1.1 404 Not Found
Content-Length: 43
Content-Type: application/json; charset=utf-8
Date: Fri, 13 Feb 2026 11:19:42 GMT
Server: backend

{
    "error": "Product to be deleted not found"
}
```
