# MinisAPI - Microservices API Gateway

## Overview
MinisAPI is a modern microservices-based API Gateway built with Go, designed to handle authentication, routing, and monitoring for distributed systems.

## Features
- Authentication & Authorization
- API Gateway routing
- Rate limiting
- Request validation
- Monitoring & Logging
- Swagger documentation

## Project Structure
```
minisapi/
├── services/           # Microservices
│   ├── auth/          # Authentication service
│   └── gateway/       # API Gateway service
├── infrastructure/    # DevOps and infrastructure
├── scripts/          # Utility scripts
└── docs/            # Documentation
```

## Prerequisites
- Go 1.21+
- Docker & Docker Compose
- Make

## Quick Start
1. Clone the repository
2. Run `make setup` to initialize the development environment
3. Run `make dev` to start the services

## Development
```bash
# Start development environment
make dev

# Run tests
make test

# Build services
make build
```

## Documentation
- [Setup Guide](docs/setup.md)
- [API Documentation](docs/api.md)
- [Architecture](docs/architecture.md)

## License
MIT License 