#!/bin/bash

# This script starts the WOW Mom App in a development environment.
# It's intended for local development and may not be secure for production use.

# Navigate to the application directory
cd wow-mom-app || {
  echo "Error: wow-mom-app directory not found. Please ensure the script is run from the project root."
  exit 1
}

# Download Go modules
echo "Downloading Go modules..."
go mod download || {
  echo "Error: Failed to download Go modules. Please check your internet connection and Go installation."
  exit 1
}

# Kill any existing server on port 30111
echo "Checking for existing processes on port 30111..."
lsof -ti:30111 | xargs -r kill -9 2>/dev/null

# Build the application
echo "Building the application..."
go build -o ../wowmom-server main.go || {
  echo "Error: Failed to build the application."
  exit 1
}

# Run the application in the background
echo "Starting WOW Mom App on port 30111..."
cd .. && ./wowmom-server &
SERVER_PID=$!

# Wait a moment for the server to start
sleep 2

# Check if the server is running
if ps -p $SERVER_PID > /dev/null; then
  echo "Server is running with PID: $SERVER_PID"
  echo "Access the application at: http://localhost:30111"
  echo ""
  echo "To stop the server, run: kill $SERVER_PID"
  echo ""
  
  # Test the connection
  echo "Testing connection..."
  if curl -s -o /dev/null -w "%{http_code}" localhost:30111 | grep -q "200"; then
    echo "✓ Server is responding correctly!"
  else
    echo "⚠ Server started but not responding as expected. Check logs for details."
  fi
  
  # Keep the script running to see server output
  echo ""
  echo "Server output:"
  echo "---"
  tail -f /dev/null & 
  wait $SERVER_PID
else
  echo "Error: Failed to start the server. Check the logs for details."
  exit 1
fi
