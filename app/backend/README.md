# Go Backend Boilerplate

A production-ready Go backend built using the Echo framework, designed with clean architecture, modular components, and modern DevOps practices.

## Architecture Overview
This project follows clean architecture principles with clear separation of layers:

- **backend/**
  - **cmd/go-boilerplate/**: Application entry point  
  - **internal/**: Core business logic  
    - **config/**: Configuration management  
    - **database/**: Database setup and migrations  
    - **handler/**: HTTP request handlers  
    - **service/**: Business logic layer  
    - **repository/**: Data access layer  
    - **model/**: Domain models  
    - **middleware/**: HTTP middleware  
    - **lib/**: Shared libraries  
    - **validation/**: Request validation  
  - **static/**: Static assets (OpenAPI specs)  
  - **templates/**: Email templates  
  - **Taskfile.yml**: Task automation  

## Features

### Core Framework
- **Echo v4**: High-performance Go web framework  
- **Clean Architecture**: Handlers → Services → Repositories → Models  
- **Dependency Injection**: Constructor-based DI for testability  

### Database
- **PostgreSQL** with pgx/v5 driver  
- **Tern** migrations for schema versioning  
- Connection pooling for performance  
- Transaction management for reliability  

### Authentication & Security
- **Clerk** integration for authentication  
- **JWT**-based access control  
- Role-based authorization  
- Rate limiting (20 req/sec per IP)  
- Security headers for XSS, CSRF, and clickjacking protection  

### Observability
- **New Relic APM** for performance monitoring  
- Structured JSON logs with **Zerolog**  
- Distributed tracing support  
- Health checks (readiness & liveness)  
- Custom business metrics  

### Background Jobs
- **Asynq**: Redis-based distributed job system  
- Priority queues (critical, default, low)  
- Retry with exponential backoff  
- Cron-style scheduling  
- Real-time job monitoring  

### Email Service
- **Resend** integration for reliable delivery  
- HTML templates for transactional emails  
- Preview mode for development  
- Batch sending support  

### API Documentation
- **OpenAPI 3.0** specification  
- **Swagger UI** for interactive API exploration  
- Auto-generation from code routes  

## Getting Started

### Prerequisites
- Go 1.24+  
- PostgreSQL 16+  
- Redis 8+  
- Task runner ([taskfile.dev](https://taskfile.dev))

### Installation
1. Install dependencies:
   ```bash
   go mod download
   ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   # Configure your environment variables
   ```

3. Run database migrations:
   ```bash
   task migrations:up
   ```

4. Start the application:
   ```bash
   task run
   ```

## Configuration
Configuration is handled through environment variables prefixed with `BOILERPLATE_`.

## Development

### Common Tasks
```bash
task help              # List all tasks
task run               # Run the app
task test              # Run tests
task migrations:new X  # Create a new migration
task migrations:up     # Apply migrations
task migrations:down   # Roll back the last migration
task tidy              # Tidy and format modules
```

### Code Structure
- **Handlers** — Parse requests, call services, return responses  
- **Services** — Contain business logic and orchestrate use cases  
- **Repositories** — Handle data persistence and queries  
- **Models** — Define domain entities and validation rules  
- **Middleware** — Handle auth, logging, rate limiting, and security  

### Testing
```bash
go test ./...
```

## Logging
Structured logging using **Zerolog**:
```go
log.Info().
  Str("user_id", userID).
  Msg("User logged in successfully")
```
Log levels: `debug`, `info`, `warn`, `error`, `fatal`

## Production Checklist
- [ ] Configure production environment variables  
- [ ] Enable SSL/TLS  
- [ ] Connect to production database  
- [ ] Set up monitoring & alerts  
- [ ] Configure log aggregation  
- [ ] Enable rate limiting  
- [ ] Establish backup strategy  
- [ ] Enable auto-scaling  
- [ ] Implement graceful shutdown  
- [ ] Integrate CI/CD pipeline  

## Performance Optimization

**Database:**  
- Connection pooling  
- Prepared statements  
- Indexed fields  
- Query optimization  

**Caching:**  
- Redis for sessions  
- In-memory caching  
- HTTP cache headers  

**Concurrency:**  
- Goroutine pools  
- Context cancellation  
- Proper synchronization  

## Security Best Practices
1. Validate and sanitize all inputs  
2. Use parameterized queries (prevent SQL injection)  
3. Apply XSS and CSRF protection  
4. Enforce rate limiting per IP/user  
5. Store secrets in environment variables  
6. Use HTTPS in production  
7. Regular dependency scanning  

## Contributing
- Follow Go idioms and clean code practices  
- Add tests for new features  
- Update documentation  
- Run linters before committing  
- Keep commits atomic and descriptive  

## License
MIT License — see LICENSE file.
