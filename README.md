# URL Shortener

A minimal, clean, and extensible URL shortener written in **Go**. This project provides REST APIs to shorten URLs, redirect them, and retrieve metrics on the most frequently shortened domains.

Github: https://github.com/harshit-0802/url-shortener

DockerHub: https://hub.docker.com/r/harshit0210/url-shortener

---

## ✨ Features

- 🔗 Shorten long URLs and get a consistent short code
- 🔁 Redirect from short code to original URL
- 📊 Metrics endpoint to return top 3 shortened domains
- 🧪 Unit test coverage for critical components
- 🔧 Configurable & decoupled architecture
- 🐳 Dockerized for easy deployment
- 📘 OpenAPI spec for API documentation

---
## REST API Endpoints

- `POST /shorten`  
  Shortens a given long URL  
  **Request Body:**
  ```json
  { "url": "https://example.com/..." }
  ```
  **Response:**
  ```json
  { "short_url": "http://localhost:8080/abc123" }
  ```

- `GET /{code}`  
  Redirects to the original long URL

- `GET /metrics`  
  Returns top 3 most shortened domains  
  **Response Example:**  
  ```json
  [
    { "domain": "youtube.com", "count": 4 },
    { "domain": "udemy.com", "count": 3 },
    { "domain": "wikipedia.org", "count": 2 }
  ]
  ```
---

## ⚙️ Getting Started

### Prerequisites

- Go 1.20+
- Docker (optional, for containerized builds)

### Running Locally

```bash
git clone https://github.com/harshit-0802/url-shortener.git
cd url-shortener

# Run directly
go run ./cmd/urlshortener
```
The service will start at http://localhost:8080
