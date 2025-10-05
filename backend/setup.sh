#!/bin/bash

# learng Backend Setup Script

set -e

echo "=================================="
echo "learng Backend Setup"
echo "=================================="
echo ""

# Check Go installation
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}')
echo "✅ Go installed: $GO_VERSION"
echo ""

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "❌ Error: go.mod not found. Please run this script from the backend directory."
    exit 1
fi

# Install dependencies
echo "📦 Installing dependencies..."
go mod download
go mod tidy
echo "✅ Dependencies installed"
echo ""

# Create .env file if it doesn't exist
if [ ! -f ".env" ]; then
    echo "📝 Creating .env file..."
    cp .env.example .env
    
    # Generate a random JWT secret
    JWT_SECRET=$(openssl rand -base64 32 2>/dev/null || cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
    
    if [ "$(uname)" == "Darwin" ]; then
        # macOS
        sed -i '' "s/your-secret-key-change-in-production/$JWT_SECRET/" .env
    else
        # Linux
        sed -i "s/your-secret-key-change-in-production/$JWT_SECRET/" .env
    fi
    
    echo "✅ .env file created with random JWT secret"
    echo ""
else
    echo "ℹ️  .env file already exists, skipping creation"
    echo ""
fi

# Create upload directories
echo "📁 Creating upload directories..."
mkdir -p uploads/images
mkdir -p uploads/audio
echo "✅ Upload directories created"
echo ""

# Optional: Install development tools
read -p "Install development tools (air for hot reload)? (y/n): " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "🔧 Installing Air (hot reload)..."
    go install github.com/cosmtrek/air@latest
    echo "✅ Development tools installed"
    echo ""
fi

# Test build
echo "🔨 Testing build..."
if go build -o tmp/learng-api ./cmd/api; then
    echo "✅ Build successful"
    rm -rf tmp
    echo ""
else
    echo "❌ Build failed. Please check for errors above."
    exit 1
fi

# Summary
echo "=================================="
echo "✅ Setup Complete!"
echo "=================================="
echo ""
echo "Next steps:"
echo "  1. Review and edit .env file if needed"
echo "  2. Run 'make run' to start the server"
echo "  3. Run 'make dev' for hot reload development"
echo "  4. Visit http://localhost:8080/health to verify"
echo ""
echo "Available commands:"
echo "  make run       - Start the server"
echo "  make dev       - Start with hot reload"
echo "  make test      - Run tests"
echo "  make help      - Show all available commands"
echo ""
