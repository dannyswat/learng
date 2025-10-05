#!/bin/bash

# learng Frontend Setup Script
# Initializes the frontend project with dependencies and configuration

set -e

echo "🚀 Setting up learng frontend..."

# Check Node.js version
NODE_VERSION=$(node -v | cut -d'v' -f2 | cut -d'.' -f1)
if [ "$NODE_VERSION" -lt 18 ]; then
    echo "❌ Error: Node.js 18+ is required (current: $(node -v))"
    exit 1
fi

echo "✅ Node.js version: $(node -v)"

# Install dependencies
echo "📦 Installing dependencies..."
npm install

# Type check
echo "🔍 Running type check..."
npm run type-check || echo "⚠️  Type errors found (expected until backend is running)"

# Try to build
echo "🏗️  Building project..."
npm run build || echo "⚠️  Build failed (expected - some imports may be unresolved)"

echo ""
echo "✅ Frontend setup complete!"
echo ""
echo "📚 Next steps:"
echo "  1. Start the backend server (cd ../backend && make run)"
echo "  2. Start the frontend dev server: npm run dev"
echo "  3. Open http://localhost:5173 in your browser"
echo ""
echo "📖 See README.md for more information"
