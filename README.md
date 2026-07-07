# sd3971-go: Product Inventory Management System

A REST API for managing products in stock, built with Go, Gin framework, and PostgreSQL.

## Features

- **CRUD Operations**: Create, Read, Update, Delete products
- **RESTful API**: Clean API endpoints for product management
- **PostgreSQL Database**: Persistent data storage with indexed queries
- **Clean Architecture**: Layered structure (handlers → services → repositories → database)
- **Docker Support**: Easy PostgreSQL setup with docker-compose

## Project Structure

```
sd3971-go/
├── main.go                          # Application entry point
├── go.mod                           # Go module definition
├── go.sum                           # Go module checksums
├── Dockerfile                       # Docker image definition
├── docker-compose.yml               # Multi-service orchestration
├── .env.example                     # Environment variables template
├── README.md                        # This file
├── config/
│   └── config.go                    # Configuration management
├── internal/
│   ├── app/
│   │   └── app.go                   # Application initialization
│   ├── routes/
│   │   └── routes.go                # Route registration
│   ├── handlers/
│   │   └── product_handler.go       # HTTP request handlers
│   ├── services/
│   │   └── product_service.go       # Business logic
│   ├── repositories/
│   │   └── product_repository.go    # Data access layer
│   ├── models/
│   │   └── product_model.go         # Data models
│   └── infrastructure/
│       └── database.go              # Database initialization
└── migrations/
    └── 001_create_products_table.up.sql # Database schema
```

## Prerequisites

- **Go 1.26+**: [Download Go](https://golang.org/dl/)
- **Docker & Docker Compose**: [Download Docker](https://www.docker.com/products/docker-desktop)
- **PostgreSQL 18+** (if running without Docker)

## Getting Started

### 1. Clone/Setup Project

```bash
cd sd3971-go
```

### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

### 3. Set Up Environment Variables

Create a `.env` file from the template:

```bash
cp .env.example .env
```

Update `.env` with your database credentials (or leave defaults for local development):

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=products_db
APP_PORT=8080
```

**Note**: When using Docker Compose, `DB_HOST` is automatically set to `postgres` (the service name in the container network).

### 4. Start Services

**Option A: Using Docker Compose (Recommended)**

Build and start both PostgreSQL and the Go app:

```bash
docker-compose up -d --build
```

Verify both services are running:

```bash
docker-compose ps
```

You should see:
- `sd3971_postgres` - Running
- `sd3971_app` - Running

Access the app at: http://localhost:8080

**Option B: Docker PostgreSQL + Local Go App**

Start only PostgreSQL:

```bash
docker-compose up -d postgres
```

Then run the Go app locally (see step 5).

### 5. Run Application

**Option A: Using Docker Compose (All-in-one)**

If you started with `docker-compose up -d --build`, the app is already running.

View logs:
```bash
docker-compose logs -f app
```

Access API: http://localhost:8080

**Option B: Local Go App with Docker PostgreSQL**

First, install Go dependencies:

```bash
go mod download
go mod tidy
```

Then start the app:

```bash
go run main.go
```

You should see:
```
Database connection established successfully
Migrations completed successfully
Application initialized successfully
Starting server on :8080
```

### 6. Test the API

#### Create a Product (POST)

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "price": 999.99,
    "quantity": 10
  }'
```

Expected Response (201 Created):
```json
{
  "id": 1,
  "name": "Laptop",
  "description": "High-performance laptop",
  "price": 999.99,
  "quantity": 10,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

#### Get All Products (GET)

```bash
curl http://localhost:8080/products
```

Expected Response (200 OK):
```json
[
  {
    "id": 1,
    "name": "Laptop",
    "description": "High-performance laptop",
    "price": 999.99,
    "quantity": 10,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
]
```

#### Get Product by ID (GET)

```bash
curl http://localhost:8080/products/1
```

Expected Response (200 OK):
```json
{
  "id": 1,
  "name": "Laptop",
  "description": "High-performance laptop",
  "price": 999.99,
  "quantity": 10,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

#### Update Product (PUT)

```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "High-Performance Laptop",
    "quantity": 15
  }'
```

Expected Response (200 OK):
```json
{
  "id": 1,
  "name": "High-Performance Laptop",
  "description": "High-performance laptop",
  "price": 999.99,
  "quantity": 15,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:35:00Z"
}
```

#### Delete Product (DELETE)

```bash
curl -X DELETE http://localhost:8080/products/1
```

Expected Response (204 No Content)

#### Health Check

```bash
curl http://localhost:8080/health
```

Expected Response (200 OK):
```json
{
  "status": "ok"
}
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/products` | Get all products |
| POST | `/products` | Create a new product |
| GET | `/products/:id` | Get product by ID |
| PUT | `/products/:id` | Update product by ID |
| DELETE | `/products/:id` | Delete product by ID |

## Database Schema

The `products` table structure:

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(12, 2) NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_products_name ON products(name);
```

## Validation Rules

### Create/Update Product

- **name**: Required, non-empty string
- **description**: Required, non-empty string
- **price**: Required, must be greater than 0
- **quantity**: Required, must be greater than or equal to 0

## Error Responses

### 400 Bad Request

```json
{
  "message": "product name is required"
}
```

### 404 Not Found

```json
{
  "message": "product not found"
}
```

### 500 Internal Server Error

```json
{
  "message": "failed to retrieve products"
}
```

## Development

### Docker Compose Commands

```bash
# Build and start both PostgreSQL and Go app
docker-compose up -d --build

# Stop all services
docker-compose down

# View logs for all services
docker-compose logs -f

# View logs for specific service
docker-compose logs -f app
docker-compose logs -f postgres

# Access PostgreSQL CLI
docker-compose exec postgres psql -U postgres -d products_db

# Rebuild app container (after code changes)
docker-compose up -d --build app

# Remove volumes (careful: deletes database)
docker-compose down -v
```

## Production Considerations

1. **Environment Configuration**: Use secure methods to manage credentials (e.g., AWS Secrets Manager, HashiCorp Vault)
2. **Database Migration**: Consider using golang-migrate for versioned migrations
3. **Logging**: Add structured logging (e.g., logrus, zap)
4. **Testing**: Add unit and integration tests
5. **Monitoring**: Add metrics and tracing (e.g., Prometheus, Jaeger)
6. **Docker Multi-Stage Build**: Optimize Docker image size for production
7. **Rate Limiting**: Add rate limiting middleware for API protection
8. **API Documentation**: Use Swagger/OpenAPI for API documentation

## Troubleshooting

### Database Connection Error

If you get "failed to connect to database" error:

1. Ensure PostgreSQL is running:
   ```bash
   docker-compose ps
   ```

2. Check database credentials in `.env` file

3. Verify PostgreSQL port (default 5432) is not in use:
   ```bash
   netstat -an | grep 5432  # Linux/Mac
   netstat -ano | findstr 5432  # Windows
   ```

### Port Already in Use

If port 8080 is already in use:

1. Change `APP_PORT` in `.env` file
2. Run: `go run main.go`

### Module Not Found

If you get "module not found" error:

1. Run: `go mod download`
2. Run: `go mod tidy`

## License

This project is for educational purposes.

## Support

For issues or questions, refer to:
- [Gin Documentation](https://github.com/gin-gonic/gin)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Go Standard Library](https://golang.org/pkg/)
