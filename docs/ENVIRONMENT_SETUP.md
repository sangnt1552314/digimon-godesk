# Environment Setup Summary

## Quick Commands

### Using Make (Recommended)
```bash
make dev          # Development with hot reloading
make prod         # Production optimized
make test         # Run tests
make run-local    # Run locally without Docker
make logs-dev     # View development logs
make logs-prod    # View production logs
make clean        # Clean up containers
```

### Using Setup Script
```bash
./scripts/setup.sh           # Interactive menu
./scripts/setup.sh dev       # Run development
./scripts/setup.sh prod      # Run production
./scripts/setup.sh test      # Run tests
```

### Manual Commands
```bash
# Development
cp .env.development .env
docker-compose -f docker-compose.dev.yml up --build

# Production  
cp .env.production .env
docker-compose -f docker-compose.prod.yml up --build

# Local run
APP_ENV=development go run ./cmd/main.go
```

## Environment Features

### Development
- **Port**: 3000
- **Hot reloading** with Air
- **Debug logging** to console and file
- **Volume mounting** for live code changes
- **Detailed error messages**

### Production
- **Port**: 8080
- **Multi-stage Docker build** for smaller images
- **File-only logging** for performance
- **Non-root user** for security
- **Health checks** included
- **Optimized binary** build

### Testing
- **Port**: 3001
- **Error-level logging** only
- **Isolated test environment**
- **Minimal resource usage**

## Configuration Files

| File | Purpose |
|------|---------|
| `.env.development` | Development settings |
| `.env.production` | Production settings |
| `.env.testing` | Testing settings |
| `docker-compose.dev.yml` | Development Docker setup |
| `docker-compose.prod.yml` | Production Docker setup |
| `Dockerfile.dev` | Development container |
| `Dockerfile.prod` | Production container |
| `Makefile` | Build automation |
| `scripts/setup.sh` | Environment setup script |

## Environment Variables

| Variable | Description | Dev | Prod | Test |
|----------|-------------|-----|------|------|
| APP_ENV | Environment name | development | production | testing |
| PORT | Server port | 3000 | 8080 | 3001 |
| DEBUG | Debug mode | true | false | false |
| LOG_LEVEL | Logging level | debug | info | error |
| LOG_FILE | Log file path | development.log | production.log | testing.log |

## Tips

1. **Always use environment-specific commands** instead of the generic docker-compose.yml
2. **Check logs** with `make logs-dev` or `make logs-prod`
3. **Clean up** regularly with `make clean`
4. **Use the setup script** for interactive environment selection
5. **Copy .env files** manually if Make is not available
