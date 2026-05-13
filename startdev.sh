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

# Run the application
echo "Starting WOW Mom App on port 30111..."
go run main.go
