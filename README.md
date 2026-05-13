# WOW Mom App

This project appears to be a web application built with Go, utilizing the Gorilla Mux router and a SQLite database. It serves static files and provides a basic backend structure.

## Getting Started

Follow these instructions to set up and run the WOW Mom App on your local machine.

### Prerequisites

*   **Go**: Ensure you have Go installed (version 1.24.4 or higher, as specified in `go.mod`). You can download it from [golang.org](https://golang.org/dl/).

### Installation

1.  **Navigate to the application directory:**
    ```bash
    cd wow-mom-app
    ```

2.  **Download Go modules:**
    This command will download all the necessary dependencies for the project.
    ```bash
    go mod download
    ```

### Running the Application

To start the application in a development environment, use the provided `startdev.sh` script:

1.  **Make the script executable (if you haven't already):**
    ```bash
    chmod +x ./startdev.sh
    ```

2.  **Run the development script:**
    ```bash
    ./startdev.sh
    ```
    This script will navigate to the `wow-mom-app` directory, download Go modules (if necessary), and then run the application. You should see output similar to:
    ```
    Downloading Go modules...
    Starting WOW Mom App on port 30111...
    ```

Alternatively, you can manually run the application from within the `wow-mom-app` directory:

```bash
cd wow-mom-app
go run main.go
```
    This will start the server. You should see output similar to:
    ```
    Server starting on port :30111
    ```

### Accessing the Application

Once the server is running, you can access the application in your web browser at:

[http://localhost:30111](http://localhost:30111)

## Database

The application uses an SQLite database named `wowmom.db`. The database schema is defined in `database/schema.sql` and is automatically applied when the application starts.
