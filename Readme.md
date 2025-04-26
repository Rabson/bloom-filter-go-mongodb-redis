# Username Check API

This project implements a scalable, fast username availability checker using a **Bloom Filter** and **MongoDB**.

---

## Strategy

User → /check-username → (Bloom → Redis → MongoDB fallback) → Fast accurate check
User → /create-user → (Check + Insert DB + Update Bloom+Key)


## Project Structure

```
username-check-api/
├── cmd/main.go           # Entry point
├── internal/
│   ├── api/              # Handlers and routes
│   ├── bloom/            # Bloom filter logic
│   ├── db/               # MongoDB connection
│   └── model/            # User model
├── load_test/            # Load testing script (k6)
├── Dockerfile            # Docker build instructions
├── docker-compose.yml    # Docker Compose for local MongoDB and API
├── swagger.yaml          # Swagger API docs
├── deployment.yaml       # Kubernetes deployment manifest
├── .dockerignore         # Docker ignore settings
├── .gitignore            # Git ignore settings
└── .github/workflows/ci-cd.yml # GitHub Actions CI/CD
```

---

## How to Run Locally

1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd username-check-api
   ```

2. Start MongoDB and the API server using Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access the API at:
   ```bash
   http://localhost:8080/check-username
   ```

## API Example

- **POST** `/check-username`
- Request Body:
  ```json
  {
    "username": "desired_username"
  }
  ```
- Response Body:
  ```json
  {
    "available": true
  }
  ```

---

## Load Testing

Run load test using [k6](https://k6.io/):

```bash
k6 run load_test/load_test.js
```

---

## Build and Run Manually

```bash
# Build
docker build -t username-check-api .

# Run
docker run -p 8080:8080 username-check-api
```

---

## Kubernetes Deployment

Deploy to a Kubernetes cluster:

```bash
kubectl apply -f deployment.yaml
```

---

## Swagger Documentation

View the API definition in the `swagger.yaml` file.
You can use [Swagger Editor](https://editor.swagger.io/) to visualize it.

---

## CI/CD Pipeline

Using **GitHub Actions**, defined in `.github/workflows/ci-cd.yml`:
- Build Docker image
- Test API
- Deploy (customizable)

---

## License

MIT License

---

## Author

Built with ❤️ by [Rabson](https://github.com/Rabson).
