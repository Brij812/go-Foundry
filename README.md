# Go Boilerplate

A production-ready monorepo template for scalable apps with Go backend and TypeScript frontend. Built with clean architecture, Turborepo, and modern tooling.

## Features

- **Monorepo**: Turborepo-powered structure  
- **Go Backend**: REST API with Echo  
- **Auth**: Clerk SDK integration  
- **Database**: PostgreSQL + migrations + pooling  
- **Jobs**: Redis async tasks with Asynq  
- **Observability**: New Relic & structured logging  
- **Email**: Resend + HTML templates  
- **Testing**: Testcontainers support  
- **API Docs**: OpenAPI/Swagger  
- **Security**: Rate limiting, CORS, secure headers, JWT  

## Project Structure

```
go-boilerplate/
├── apps/backend/      # Go backend
├── packages/          # Frontend packages
├── package.json       # Monorepo config
├── turbo.json         # Turborepo config
└── README.md          # This file
```

## Quick Start

### Prerequisites
- Go ≥ 1.24  
- Node.js 22+ & Bun  
- PostgreSQL ≥ 16  
- Redis ≥ 8  

### Installation
```bash
git clone https://github.com/sriniously/go-boilerplate.git
cd go-boilerplate

# Frontend deps
bun install

# Backend deps
cd apps/backend
go mod download

# Env setup
cp .env.example .env
# Edit .env
```

Start DB & Redis, then run migrations:
```bash
task migrations:up
```

Start dev server:
```bash
# Full stack
bun dev

# Backend only
cd apps/backend
task run
```

API runs at `http://localhost:8080`.

## Development

### Commands
```bash
# Backend
task help
task run
task migrations:new
task migrations:up
task test
task tidy

# Frontend
bun dev
bun build
bun lint
```

### Environment Variables
Prefix: `BOILERPLATE_`  
Key vars:
- `DATABASE_*` – PostgreSQL  
- `SERVER_*` – Server config  
- `AUTH_*` – Authentication  
- `REDIS_*` – Redis  
- `EMAIL_*` – Email service  
- `OBSERVABILITY_*` – Monitoring  

See `apps/backend/.env.example` for full list.

## Architecture
- **Handlers**: HTTP interface  
- **Services**: Business logic  
- **Repositories**: Data access  
- **Models**: Domain entities  
- **Infrastructure**: DB, cache, email  

## Testing
```bash
# Backend
cd apps/backend
go test ./...
go test -cover ./...
go test -tags=integration ./...
```

## Production
- Use env-specific config  
- Enable production logging  
- Configure DB pooling  
- Setup monitoring & alerts  
- Use reverse proxy (nginx, Caddy)  
- Enable rate limiting & security headers  
- Configure CORS  

## Contributing
1. Fork & create branch (`feature/...`)  
2. Commit changes  
3. Push branch  
4. Open Pull Request  

## License
MIT License — see LICENSE file.
