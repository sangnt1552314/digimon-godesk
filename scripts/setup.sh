#!/bin/bash

# Environment Setup Script for Digimon GoDesk
# This script helps you quickly set up and run the application in different environments

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Functions
print_header() {
    echo -e "${BLUE}🦕 Digimon GoDesk Environment Setup${NC}"
    echo "======================================"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

check_requirements() {
    echo "Checking requirements..."
    
    # Check Go
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed. Please install Go 1.23 or higher."
        exit 1
    fi
    print_success "Go is installed: $(go version)"
    
    # Check Docker (optional)
    if command -v docker &> /dev/null; then
        print_success "Docker is installed: $(docker --version)"
        DOCKER_AVAILABLE=true
    else
        print_warning "Docker is not installed. You can still run locally."
        DOCKER_AVAILABLE=false
    fi
    
    # Check Make (optional)
    if command -v make &> /dev/null; then
        print_success "Make is available"
        MAKE_AVAILABLE=true
    else
        print_warning "Make is not available. You'll need to run commands manually."
        MAKE_AVAILABLE=false
    fi
}

setup_environment() {
    local env=$1
    
    echo "Setting up $env environment..."
    
    # Copy environment file
    if [ -f ".env.$env" ]; then
        cp ".env.$env" ".env"
        print_success "Environment file copied: .env.$env -> .env"
    else
        print_error "Environment file .env.$env not found!"
        exit 1
    fi
    
    # Create logs directory
    mkdir -p storage/logs
    print_success "Logs directory created"
}

run_development() {
    setup_environment "development"
    
    if [ "$MAKE_AVAILABLE" = true ]; then
        echo "Running with Make..."
        make dev
    elif [ "$DOCKER_AVAILABLE" = true ]; then
        echo "Running with Docker Compose..."
        docker-compose -f docker-compose.dev.yml up --build
    else
        echo "Running locally..."
        go run ./cmd/main.go
    fi
}

run_production() {
    setup_environment "production"
    
    if [ "$MAKE_AVAILABLE" = true ]; then
        echo "Running with Make..."
        make prod
    elif [ "$DOCKER_AVAILABLE" = true ]; then
        echo "Running with Docker Compose..."
        docker-compose -f docker-compose.prod.yml up --build
    else
        echo "Building and running locally..."
        go build -o tmp/main ./cmd/main.go
        ./tmp/main
    fi
}

run_tests() {
    setup_environment "testing"
    
    if [ "$MAKE_AVAILABLE" = true ]; then
        make test
    else
        echo "Running tests..."
        APP_ENV=testing go test ./...
    fi
}

show_menu() {
    echo ""
    echo "What would you like to do?"
    echo "1) Run in Development mode"
    echo "2) Run in Production mode"
    echo "3) Run Tests"
    echo "4) Just setup environment files"
    echo "5) Exit"
    echo ""
    read -p "Enter your choice (1-5): " choice
    
    case $choice in
        1)
            run_development
            ;;
        2)
            run_production
            ;;
        3)
            run_tests
            ;;
        4)
            echo "Available environments:"
            echo "- development"
            echo "- production"
            echo "- testing"
            echo ""
            read -p "Which environment to setup? " env
            setup_environment "$env"
            print_success "Environment $env is now active"
            ;;
        5)
            echo "Goodbye! 👋"
            exit 0
            ;;
        *)
            print_error "Invalid choice. Please try again."
            show_menu
            ;;
    esac
}

# Main execution
main() {
    print_header
    check_requirements
    
    # If arguments provided, run directly
    if [ $# -gt 0 ]; then
        case $1 in
            "dev"|"development")
                run_development
                ;;
            "prod"|"production")
                run_production
                ;;
            "test"|"testing")
                run_tests
                ;;
            *)
                print_error "Unknown environment: $1"
                echo "Usage: $0 [dev|prod|test]"
                exit 1
                ;;
        esac
    else
        show_menu
    fi
}

# Run main function with all arguments
main "$@"
