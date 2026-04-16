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

---

# Phase 3: HTTP Status Codes, Error Handling & Validation

---

# 1️⃣ HTTP Status Codes

## ✅ Success Codes

| Status Code | Meaning | Used For |
|------------|----------|----------|
| 200 OK | Request successful | GET, PUT, DELETE |
| 201 Created | Resource successfully created | POST |
| 204 No Content | Resource successfully deleted (optional alternative to 200) | DELETE |

### POST Additional Requirement
- Must return `Location` header:
```
Location: /products/{id}
```

---

## ❌ Client Error Codes

| Status Code | Meaning | When Returned |
|------------|----------|---------------|
| 400 Bad Request | Invalid request format | Malformed JSON, missing required fields |
| 404 Not Found | Resource does not exist | Product not found |
| 422 Unprocessable Entity | Valid JSON but business rule fails | Validation errors |

---

## 🔥 Server Error Codes

| Status Code | Meaning |
|------------|----------|
| 500 Internal Server Error | Unexpected server-side error |

---

# 2️⃣ Standardized Error Response Format

All error responses must follow this structure:

```json
{
  "error": {
    "code": "PRODUCT_NOT_FOUND",
    "message": "Product with ID PROD-12345678 does not exist",
    "details": {}
  }
}
```

## Error Object Fields

| Field | Description |
|-------|------------|
| code | Machine-readable error identifier |
| message | Human-readable explanation |
| details | Field-level validation errors (if applicable) |

---

# 3️⃣ Defined Error Codes

| Error Code | HTTP Status | Meaning |
|------------|------------|----------|
| PRODUCT_NOT_FOUND | 404 | Product with given ID does not exist |
| INVALID_INPUT | 400 | Malformed JSON or invalid request format |
| MISSING_REQUIRED_FIELD | 400 | Required field not provided |
| VALIDATION_FAILED | 422 | Business validation rule failed |
| INTERNAL_ERROR | 500 | Unexpected system error |

---

# 4️⃣ Input Validation Rules

Validation applies to **POST** and **PUT** operations.

---

## 📌 Name

- Required
- Must be non-empty
- Maximum 200 characters

| Condition | Response |
|------------|-----------|
| Missing | 400 Bad Request |
| Empty | 400 Bad Request |
| Too long (>200) | 422 Unprocessable Entity |

---

## 📌 Price

- Required
- Must be a valid number
- Must be greater than 0
- Maximum 2 decimal places

| Condition | Response |
|------------|-----------|
| Missing | 400 Bad Request |
| Not a number | 400 Bad Request |
| ≤ 0 | 422 Unprocessable Entity |
| More than 2 decimal places | 422 Unprocessable Entity |

---

## 📌 Category ID

- Required
- Must be non-empty
- Must match pattern: `CAT-xxxxxxxx` (optional format validation)

| Condition | Response |
|------------|-----------|
| Missing | 400 Bad Request |
| Empty | 400 Bad Request |
| Invalid format (if validated) | 422 Unprocessable Entity |

---

## 📌 Description

- Optional
- Maximum 500 characters

| Condition | Response |
|------------|-----------|
| Too long (>500) | 422 Unprocessable Entity |

---

# 5️⃣ Validation Error Response Example

```json
{
  "error": {
    "code": "VALIDATION_FAILED",
    "message": "Input validation failed",
    "details": {
      "name": "Name is required and cannot be empty",
      "price": "Price must be greater than 0"
    }
  }
}
```

---

# 6️⃣ Response Headers

| Header | Applies To | Requirement |
|--------|------------|------------|
| Content-Type: application/json | All responses | Required |
| Location: /products/{id} | POST success | Required |

---

# 7️⃣ Status Codes Per Endpoint

---

## GET /products

| Scenario | Status |
|----------|--------|
| Success | 200 OK |

---

## GET /products/{id}

| Scenario | Status |
|----------|--------|
| Success | 200 OK |
| Product not found | 404 Not Found |

---

## POST /products

| Scenario | Status |
|----------|--------|
| Created successfully | 201 Created |
| Malformed JSON | 400 Bad Request |
| Missing required fields | 400 Bad Request |
| Validation failed | 422 Unprocessable Entity |
| Unexpected error | 500 Internal Server Error |

---

## PUT /products/{id}

| Scenario | Status |
|----------|--------|
| Updated successfully | 200 OK |
| Product not found | 404 Not Found |
| Malformed JSON | 400 Bad Request |
| Missing required fields | 400 Bad Request |
| Validation failed | 422 Unprocessable Entity |
| Unexpected error | 500 Internal Server Error |

---

## DELETE /products/{id}

| Scenario | Status |
|----------|--------|
| Deleted successfully | 200 OK or 204 No Content |
| Product not found | 404 Not Found |
| Unexpected error | 500 Internal Server Error |

---


# Phase 3: HTTP Status Codes, Error Handling & Validation

# Running with Containers

## Build the Image

Build the application container locally:
```bash
podman build -t ecommerce-api:dev .
```

---

## Run with Compose

Start the API and PostgreSQL services:
```bash
podman-compose up --build
```

This will:
- Build the API image
- Start the database service
- Expose the API on `http://localhost:8080`

---

## Verify the API

Check the health endpoint:
```bash
curl http://localhost:8080/health
```

If successful, it should return HTTP 200.

---

# Environment Variables

The application is configured using the following environment variables:

| Variable | Description |
|----------|------------|
| PORT | Port on which the API runs |
| DB_HOST | Database host |
| DB_PORT | Database port |
| DB_USER | Database username |
| DB_PASSWORD | Database password |
| DB_NAME | Database name |

These variables are defined in `podman-compose.yml` for local development.

---

# CI Pipeline

The CI pipeline is located at:
```
.github/workflows/ci.yml
```

## Triggers

The pipeline runs:
- On push to `main`
- On pull request to `main`

---

## Pipeline Stages

### 1. Lint

Runs static code analysis using:
- `golangci-lint` (Go projects)
The job fails if lint errors are detected.

---

### 2. Test

- Sets up test environment
- Runs unit tests
- Runs integration tests
- Generates coverage report

Minimum required coverage: **70%**

The job fails if:
- Any test fails
- Coverage is below 70%

---

### 3. Build

- Builds the container image
- Verifies the Dockerfile builds successfully

---
