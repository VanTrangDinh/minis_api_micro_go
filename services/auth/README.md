# Authentication Service

This is a microservice for handling authentication and authorization in the system.

## Features

- User registration and login
- JWT-based authentication
- Role-based access control (RBAC)
- Two-factor authentication
- Email verification
- Password reset
- Rate limiting
- Request logging
- Health check

## Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Redis
- SMTP server (for email)

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
# Server
SERVER_PORT=8080
SERVER_READ_TIMEOUT=10s
SERVER_WRITE_TIMEOUT=10s

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=auth_db
DB_SSL_MODE=disable

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT
JWT_SECRET_KEY=your-secret-key
JWT_ACCESS_TOKEN_TTL=15m
JWT_REFRESH_TOKEN_TTL=24h

# Email
EMAIL_HOST=smtp.gmail.com
EMAIL_PORT=587
EMAIL_USERNAME=your-email@gmail.com
EMAIL_PASSWORD=your-password
EMAIL_FROM=your-email@gmail.com
```

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Run migrations:
   ```bash
   go run cmd/migrate/main.go
   ```
4. Start the service:
   ```bash
   go run cmd/api/main.go
   ```

## Docker

Build the image:
```bash
docker build -t auth-service .
```

Run the container:
```bash
docker run -p 8080:8080 --env-file .env auth-service
```

## API Endpoints

### Public Routes

- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user
- `POST /api/v1/auth/refresh` - Refresh access token
- `POST /api/v1/auth/reset-password` - Request password reset
- `POST /api/v1/auth/verify-email` - Verify email address

### Protected Routes

- `POST /api/v1/auth/logout` - Logout user
- `POST /api/v1/auth/change-password` - Change password
- `POST /api/v1/auth/enable-2fa` - Enable two-factor authentication
- `POST /api/v1/auth/disable-2fa` - Disable two-factor authentication
- `POST /api/v1/auth/verify-2fa` - Verify two-factor code

### Health Check

- `GET /health` - Check service health

## Testing

Run tests:
```bash
go test ./...
```

## License

MIT 