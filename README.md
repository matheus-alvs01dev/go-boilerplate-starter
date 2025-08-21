# Go Boilerplate Starter

A modern Go boilerplate with clean architecture, database migrations, and development tools configured out of the box.

## 🚀 Features

- **Clean Architecture** - Hexagonal architecture with clear separation of concerns
- **Database Ready** - PostgreSQL integration with migrations using Goose
- **Type-safe SQL** - SQLC for generating type-safe Go code from SQL
- **HTTP Router** - Chi router with middleware support
- **Configuration** - Viper for configuration management
- **Logging** - Structured logging with Zap
- **Hot Reload** - Air for development hot reloading
- **Docker Support** - Development services with Docker Compose
- **Development Tools** - Mise for tool and task management
- **Code Quality** - Linting with golangci-lint
- **API Testing** - Bruno collection included
- **Monitoring** - Prometheus and Grafana for observability
- **Message Queue** - Kafka for event streaming
- **Caching** - Redis for caching

## 📁 Project Structure

```
.
├── cmd/
│   └── server/              # Application entrypoints
├── config/                  # Configuration files
├── internal/
│   ├── adapters/
│   │   └── db/             # Database adapters and migrations
│   ├── app/
│   │   ├── api/            # HTTP handlers and routes
│   │   ├── di/             # Dependency injection
│   │   └── setup/          # Application setup
│   └── domain/             # Bound Contexts entities, uc and any
├── pkg/                    # Public packages
├── dev/                    # Development tools and configs
└── tmp/                    # Temporary build files
```

## 🛠️ Prerequisites

- [Mise](https://mise.jdx.dev/) - Tool and runtime manager
- Docker and Docker Compose

## ⚡ Quick Start

### 1. Install Mise

```bash
# Install Mise
curl https://mise.run | sh

# Add to your shell profile
echo 'eval "$(mise activate zsh)"' >> ~/.zshrc    # for zsh
echo 'eval "$(mise activate bash)"' >> ~/.bashrc  # for bash

# Reload your shell
source ~/.zshrc  # or source ~/.bashrc
```

### 2. Setup Project

```bash
# Clone and enter the project
git clone <your-repo>
cd go-boilerplate-starter

# Trust Mise configuration
mise trust

# Install all tools and dependencies
mise install
mise run setup
```

### 3. Configure Environment

```bash
# Copy example environment file
cp .env.example .env

# Edit .env with your database configuration
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=postgres
# DB_NAME=postgres
# API_PORT=8080
```

### 4. Start Development

```bash
# Start development services (database, redis, kafka, monitoring)
mise run docker-up

# Run migrations
mise run migration-up

# Start development server with hot reload
mise run dev
```

## 🔧 Available Commands

Use `mise tasks` to see all available tasks:

### Development

- `mise run server` - Start the server
- `mise run dev` - Start with hot reload (recommended for development)
- `mise run build` - Build the application
- `mise run clean` - Clean build artifacts

### Testing

- `mise run test` - Run tests with coverage
- `mise run test-watch` - Run tests in watch mode

### Database

- `mise run migration-create name=<name>` - Create new migration
- `mise run migration-up` - Run pending migrations
- `mise run migration-down` - Rollback last migration
- `mise run migration-status` - Check migration status
- `mise run sqlc-generate` - Generate type-safe Go code from SQL

### Docker

- `mise run docker-up` - Start Docker services
- `mise run docker-down` - Stop Docker services
- `mise run docker-logs` - View Docker logs

### Code Quality

- `mise run lint` - Run linter

### Setup

- `mise run install-deps` - Install Go development dependencies
- `mise run setup` - Complete project setup (deps + database)

## 🗃️ Database

### Migrations

Migrations are located in `internal/adapters/db/migrations/` and managed with Goose:

```bash
# Create a new migration
mise run migration-create name=create_users_table

# Run migrations
mise run migration-up

# Check status
mise run migration-status
```

### SQLC

SQL queries are in `internal/adapters/db/sqlc/queries/` and automatically generate type-safe Go code:

```bash
# Generate Go code from SQL
mise run sqlc-generate
```

## 🏗️ Architecture

This boilerplate follows **Hexagonal Architecture** (Ports and Adapters):

- **`cmd/`** - Application entry points
- **`internal/domain/`** - Core business logic (entities, use cases)
- **`internal/adapters/`** - External adapters (database, HTTP, etc.)
- **`internal/app/`** - Application layer (DI, setup, API)
- **`pkg/`** - Reusable packages
- **`config/`** - Configuration management

## 🧪 Testing

```bash
# Run all tests
mise run test

# Run tests in watch mode
mise run test-watch

# Run tests with verbose output
go test -v ./...
```

## 🐳 Docker Services

The project includes several development services via Docker Compose:

### Available Services

- **PostgreSQL** - Primary database (port 5432)
- **Redis** - Caching and session storage (port 6379)
- **Kafka** - Message streaming (port 9092)
- **Zookeeper** - Kafka coordination (port 2181)
- **Prometheus** - Metrics collection (port 9090)
- **Grafana** - Monitoring dashboards (port 3000, admin/admin)

### Commands

```bash
# Start all services
mise run docker-up

# View logs
mise run docker-logs

# Stop services
mise run docker-down
```

### Service URLs

- Grafana Dashboard: http://localhost:3000 (admin/admin)
- Prometheus: http://localhost:9090
- PostgreSQL: localhost:5432
- Redis: localhost:6379
- Kafka: localhost:9092

## 📝 API Documentation

API testing collection is available in `dev/bruno/`. Import the collection into [Bruno](https://www.usebruno.com/) for interactive API testing.

## ⚙️ Configuration

Configuration is managed with Viper. See `config/config.yaml` for available options.

Environment variables:

- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database username (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: postgres)
- `DATABASE_URL` - PostgreSQL connection string (auto-generated from above)
- `API_PORT` - Server port (default: 8080)
- `LOG_LEVEL` - Logging level (debug, info, warn, error)

## 🔧 Development Tools

This project uses **Mise** for managing development tools and tasks:

- **Go** - Programming language
- **SQLC** - Type-safe SQL code generation
- **Goose** - Database migrations
- **Air** - Hot reload for development
- **golangci-lint** - Go linter
- **mockgen** - Mock generation for testing

### Local Configuration

Create `.mise.local.toml` for personal overrides:

```toml
[env]
DATABASE_URL = "postgres://user:pass@localhost:5432/mydb"
API_PORT = "3000"

[tasks.my-task]
description = "My personal task"
run = "echo 'Hello World'"
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run linting: `mise run lint`
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Chi](https://github.com/go-chi/chi) - HTTP router
- [SQLC](https://sqlc.dev/) - SQL code generation
- [Goose](https://github.com/pressly/goose) - Database migrations
- [Viper](https://github.com/spf13/viper) - Configuration management
- [Zap](https://github.com/uber-go/zap) - Logging
- [Air](https://github.com/air-verse/air) - Hot reload
- [Mise](https://mise.jdx.dev/) - Tool management
- [Prometheus](https://prometheus.io/) - Monitoring
- [Grafana](https://grafana.com/) - Dashboards
