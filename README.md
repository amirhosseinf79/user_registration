# OTP-Based Authentication Service

A robust Golang backend service implementing OTP-based login and registration with comprehensive user management features.

## Features

- **OTP Authentication**: Phone number-based registration and login
- **Rate Limiting**: 3 OTP requests per phone number within 10 minutes
- **JWT Authentication**: Secure token-based authentication
- **User Management**: RESTful APIs with pagination and search
- **Clean Architecture**: Repository pattern with SOLID principles
- **Swagger Documentation**: Complete API documentation
- **Docker Support**: Containerized deployment ready

## Architecture

- **Backend**: Golang with Fiber framework
- **Primary Database**: PostgreSQL for user data persistence
- **Cache Layer**: Redis for OTP storage and rate limiting
- **Authentication**: JWT Bearer tokens
- **Documentation**: Swagger/OpenAPI

## Prerequisites

- Docker and Docker Compose
- Git

## Docker Setup

### Quick Start

1. **Clone the repository**
   ```bash
   git clone git@github.com:amirhosseinf79/user_registration.git
   cd user_registration
   ```

2. **Create environment file**
   ```bash
   cp .env.example .env
   ```

3. **Run with Docker Compose**
   ```bash
   docker-compose up --build
   ```

4. **Access the application**
   - API: http://localhost:8080
   - Swagger Docs: http://localhost:8080/swagger/index.html

### Environment Variables

Create a `.env` file with:

```env
# Server configuration
PORT=8080

# Application secret key
SECRET=your_secret_key_here

# Database connection string (GORM)
SQLDB="host=db user=postgres password=postgres dbname=user_otp port=5432 sslmode=disable TimeZone=Asia/Tehran"

# Redis configuration
RedisServer=localhost:6379
RedisPass=your_redis_password_here

# Debug mode
DEBUG=true
```

### Docker Services

```yaml
services:
  db:
    image: postgres:17
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: user_otp
    volumes:
      - db-data:/var/lib/postgresql/data

  redis:
    image: redis:7
    volumes:
      - redis-data:/data

  app:
    build:
      context: .
      dockerfile: cmd/Dockerfile
    ports:
      - "8080:8080"
    depends_on:
      - db
      - redis
    environment:
      - PORT=8080
      - SECRET=24vm89v5y7q-x,m349ci-143-v5um120-5v27n45-1237cn4
      - SQLDB=host=db user=postgres password=postgres dbname=user_otp port=5432 sslmode=disable TimeZone=Asia/Tehran
      - RedisServer=redis:6379
      - RedisPass=
      - DEBUG=false

volumes:
  db-data:
  redis-data:
```

### Useful Commands

```bash
# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Rebuild and restart
docker-compose up --build

# Clean up (removes data!)
docker-compose down -v
```

## Database Choice Justification

### Why PostgreSQL + Redis?

#### PostgreSQL (Primary Database)
- **ACID Compliance**: Ensures data integrity for user registration
- **Advanced Features**: JSON support, UUID generation, efficient indexing
- **Scalability**: Excellent performance for read-heavy workloads and pagination
- **Search Capabilities**: Full-text search and pattern matching for user queries
- **Production Ready**: Mature ecosystem with robust backup and monitoring tools

#### Redis (Cache Layer)
- **Performance**: Sub-millisecond response times for OTP verification
- **Built-in TTL**: Automatic OTP expiration (2 minutes) without cleanup jobs
- **Atomic Operations**: Thread-safe rate limiting and attempt counting
- **Memory Efficiency**: Optimized for temporary data storage

### Data Flow
```
User Request → PostgreSQL (User Data) + Redis (OTP/Rate Limiting)
```

### Alternative Considerations
- **MongoDB**: Rejected due to eventual consistency model
- **MySQL**: PostgreSQL offers better JSON support and concurrency
- **SQLite**: Not suitable for production concurrent access
- **In-Memory Only**: Data loss on restart, not suitable for user data

## API Endpoints

- `POST /auth/send-otp` - Request OTP
- `POST /auth/verify-otp` - Verify OTP and login/register
- `GET /user/all` - Get users (with pagination and search)
- `GET /user/:userID` - Get user by ID
- `GET /profile/update` - Update user Profile

## Request: POST /auth/send-otp
```json
{
    "phoneNumber": "09334455678"
}
```

## Response:
```json
{
  "message": "ok"
}
```

## Request: POST /auth/verify-otp
```json
{
  "phone": "09334455678",
  "code": "123456"
}
```

## Response:
```
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
  "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```

## Development

### Local Development
```bash
# Install dependencies
go mod download

# Run locally (requires PostgreSQL and Redis)
go run main.go
```

## Security Features

- JWT token authentication
- Rate limiting to prevent OTP spam
- Phone number validation
- Secure OTP generation (6-digit random)
- Password-less authentication

## Documentation

Complete API documentation + Example API requests & responses is available at `/swagger/index.html` when the application is running.
