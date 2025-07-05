# 🦕 Digimon GoDesk

A Digimon information webapp built with Go to help me study Golang while exploring the digital world of [Digimon](https://en.wikipedia.org/wiki/Digimon)! 

## 🎯 Purpose

This project serves as my hands-on learning experience with Golang. What better way to learn a new language than building something about Digimon? After all, both involve evolution and transformation! 🚀

## 😅 A Funny Confession

**Full transparency:** I don't know CSS, and GitHub Copilot has been my styling savior throughout this project. Every beautiful button, every aligned div, every responsive layout - all thanks to AI assistance. If you see any decent styling, know that it wasn't me! 🎨✨

## 🌐 Data Source

This project utilizes the fantastic [Digi-API](https://digi-api.com/) to fetch Digimon information. Huge thanks to the maintainers for providing such a comprehensive and well-structured API for Digimon data!

## 🏗️ Project Structure

```
digimon-godesk/
├── cmd/                    # Application entry point
│   └── main.go
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── handler/          # HTTP handlers
│   │   ├── api/         # API endpoints
│   │   ├── middleware/  # HTTP middleware
│   │   └── web/         # Web page handlers
│   ├── logger/          # Logging utilities
│   ├── models/          # Data models
│   ├── services/        # Business logic
│   └── utils/           # Utility functions
├── templates/            # HTML templates
├── assets/              # Static assets (CSS, JS, images)
├── storage/             # Application storage
└── docker-compose.yml   # Docker setup
```

## 🚀 Getting Started

### Prerequisites

- **Go 1.23 or higher** - Required for local development
- **Docker & Docker Compose** - Required for containerized environments
- **Make** - Optional but recommended for simplified commands

### Quick Setup

1. **Clone the repository:**
```bash
git clone https://github.com/sangnt1552314/digimon-godesk.git
cd digimon-godesk
```

2. **Install dependencies:**
```bash
go mod download
# or use make command
make deps
```

### 🔥 Development Environment

Perfect for coding, debugging, and testing with hot reloading enabled.

**Using Make (Recommended):**
```bash
# Start development server with hot reloading
make dev

# Run in background
make dev-bg

# View development logs
make logs-dev

# Stop development server
make stop-dev
```

**Manual Setup:**
```bash
# Copy development configuration
cp .env.development .env

# Option 1: Docker with hot reloading (Recommended)
docker-compose -f docker-compose.dev.yml up --build

# Option 2: Run locally without Docker
APP_ENV=development go run ./cmd/main.go
```

**Development Features:**
- 🔄 **Hot reloading** - Code changes automatically restart the server
- 🐛 **Debug mode** - Detailed error messages and debug routes
- 📝 **Verbose logging** - Debug level logging for troubleshooting
- 🌐 **Accessible at:** `http://localhost:3000`

### 🚀 Production Environment

Optimized for performance, security, and stability.

**Using Make (Recommended):**
```bash
# Start production server
make prod

# Run in background
make prod-bg

# View production logs
make logs-prod

# Stop production server
make stop-prod
```

**Manual Setup:**
```bash
# Copy production configuration
cp .env.production .env

# Option 1: Docker (Recommended)
docker-compose -f docker-compose.prod.yml up --build

# Option 2: Build and run locally
APP_ENV=production go build -o main ./cmd/main.go
./main
```

**Production Features:**
- ⚡ **Optimized build** - Multi-stage Docker build for minimal image size
- 🔒 **Security hardened** - Debug routes disabled, minimal logging
- 💪 **Health checks** - Built-in health monitoring
- 🌐 **Accessible at:** `http://localhost:8080`

## 🌍 Environment Management

This project supports multiple environments with different configurations optimized for specific use cases.

### 📋 Available Environments

| Environment | Purpose | Port | Hot Reload | Debug | Log Level |
|-------------|---------|------|------------|--------|-----------|
| **Development** | Coding & debugging | 3000 | ✅ Yes | ✅ Enabled | Debug |
| **Production** | Live deployment | 8080 | ❌ No | ❌ Disabled | Info |
| **Testing** | Running tests | 3001 | ❌ No | ❌ Disabled | Error |

### 🛠️ Make Commands (Recommended)

The easiest way to manage the application across different environments:

```bash
# 🔧 Development Commands
make dev          # Start development server (foreground)
make dev-bg       # Start development server (background)
make logs-dev     # View development logs
make stop-dev     # Stop development server
make restart-dev  # Restart development server

# 🚀 Production Commands  
make prod         # Start production server (foreground)
make prod-bg      # Start production server (background)
make logs-prod    # View production logs
make stop-prod    # Stop production server
make restart-prod # Restart production server

# 🧪 Testing & Utilities
make test         # Run tests
make build        # Build application locally
make run-local    # Run locally without Docker
make clean        # Clean up containers and images
make status       # View container status
make deps         # Install dependencies
make help         # Show all available commands
```

### 📁 Environment Configuration Files

- **`.env.development`** - Development settings (hot reload, debug mode)
- **`.env.production`** - Production settings (optimized, secure)
- **`.env.testing`** - Testing settings (isolated, minimal logging)
- **`.env.example`** - Template for creating custom `.env` files

### 🔧 Manual Environment Setup

#### Development Environment Setup

```bash
# Method 1: Using Make (Recommended)
make dev

# Method 2: Manual Docker setup
cp .env.development .env
docker-compose -f docker-compose.dev.yml up --build

# Method 3: Local development without Docker
cp .env.development .env
APP_ENV=development go run ./cmd/main.go
```

**Development Environment Features:**
- 🔄 **Air hot reloading** - Automatically restarts on code changes
- 🐛 **Debug routes enabled** - Access to `/debug/` endpoints
- 📝 **Verbose logging** - Detailed logs for development
- 🌐 **Development port** - Runs on `http://localhost:3000`

#### Production Environment Setup

```bash
# Method 1: Using Make (Recommended)
make prod

# Method 2: Manual Docker setup
cp .env.production .env
docker-compose -f docker-compose.prod.yml up --build

# Method 3: Manual build and run
cp .env.production .env
APP_ENV=production go build -o main ./cmd/main.go
./main
```

**Production Environment Features:**
- ⚡ **Multi-stage Docker build** - Optimized for size and performance
- 🔒 **Security hardened** - Debug routes disabled, minimal error exposure
- 💪 **Health checks** - Container health monitoring
- 🌐 **Production port** - Runs on `http://localhost:8080`

### 🔐 Environment Variables Reference

| Variable | Development | Production | Testing | Description |
|----------|-------------|------------|---------|-------------|
| `APP_ENV` | development | production | testing | Application environment |
| `PORT` | 3000 | 8080 | 3001 | Server port |
| `DEBUG` | true | false | false | Enable debug mode |
| `LOG_LEVEL` | debug | info | error | Logging level |
| `LOG_FILE` | development.log | production.log | testing.log | Log file location |
| `HOT_RELOAD` | true | false | false | Enable hot reloading |
| `ENABLE_DEBUG_ROUTES` | true | false | false | Enable debug endpoints |

### 🐳 Docker Configuration Files

- **`docker-compose.yml`** - Default configuration (development)
- **`docker-compose.dev.yml`** - Development with hot reloading and volume mounts
- **`docker-compose.prod.yml`** - Production with health checks and optimizations

### 📦 Dockerfile Options

- **`Dockerfile`** - Default development setup
- **`Dockerfile.dev`** - Development with Air hot reloading
- **`Dockerfile.prod`** - Multi-stage production build (optimized)

### 🔍 Monitoring & Logs

```bash
# View real-time logs
make logs-dev     # Development logs
make logs-prod    # Production logs

# View application logs (file-based)
tail -f storage/logs/development.log  # Development
tail -f storage/logs/production.log   # Production

# Container status
make status       # View all containers
docker ps -a      # All Docker containers
```

### 🧹 Cleanup & Maintenance

```bash
# Stop and clean everything
make clean        # Remove containers, volumes, and prune system

# Individual cleanup
make stop-dev     # Stop development containers
make stop-prod    # Stop production containers

# Rebuild from scratch
make clean && make dev    # Clean and restart development
make clean && make prod   # Clean and restart production
```

## 🔧 Features

- 📱 Responsive web interface (thanks to AI styling!)
- 🔍 Search and browse Digimon
- 📊 Display detailed Digimon information
- 🌐 RESTful API endpoints
- 📝 Comprehensive logging
- 🐳 Docker support

## 🧠 What I'm Learning

Through this project, I'm exploring:

- Go web frameworks and HTTP handling
- Project structure and organization in Go
- API integration and HTTP clients
- Template rendering and web development
- Docker containerization
- Middleware implementation
- Logging and error handling
- And most importantly... how to make things look good without knowing CSS! 😂

## 🦖 About Digimon

[Digimon](https://en.wikipedia.org/wiki/Digimon) (Digital Monsters) is a Japanese media franchise encompassing virtual pet toys, anime, manga, video games, films and a trading card game. The franchise focuses on creatures called "Digimon" who inhabit a "Digital World," parallel to the real world.

Just like how Digimon evolve and grow stronger, I hope my Go skills evolve through building this project! 💪

## 🤝 Contributing

Feel free to contribute! Whether it's Go code improvements, CSS fixes (because I definitely need help there), or new features - all contributions are welcome!

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

- [Digi-API](https://digi-api.com/) - For providing the comprehensive Digimon data API
- GitHub Copilot - For being my CSS styling superhero 🦸‍♂️
- The Digimon community - For keeping the digital spirit alive
- The Go community - For creating such an awesome language to learn!

---

*"The courage to take the first step is what separates the dreamer from the doer."* - Tai Kamiya

**Happy coding and digital evolution!** 🌟

### 🔧 Troubleshooting

#### Common Issues and Solutions

**Port Already in Use:**
```bash
# Check what's using the port
lsof -i :3000  # Development
lsof -i :8080  # Production

# Kill the process
kill -9 <PID>

# Or use different ports by editing .env files
```

**Docker Issues:**
```bash
# Clean Docker cache
docker system prune -a

# Remove all containers and start fresh
make clean && make dev
```

**Hot Reload Not Working:**
```bash
# Ensure you're using the development environment
make dev

# Check if Air is installed in the container
docker exec -it digimon-godesk-dev air -v
```

**Cannot Access Application:**
```bash
# Check container status
make status

# Check logs for errors
make logs-dev    # Development
make logs-prod   # Production

# Verify port mapping
docker ps | grep digimon-godesk
```

**Environment Variables Not Loading:**
```bash
# Ensure .env file exists and is properly formatted
cat .env

# Check if environment variables are set correctly
docker exec -it digimon-godesk-dev env | grep APP_ENV
```