# Golang Template Project

This is a Golang project template or boilerplate code designed to kickstart the development of RESTful APIs using the Gin framework. The template includes basic endpoints for adding and retrieving user information, integrates the Zap logger for efficient logging, utilizes middleware to log incoming requests, and leverages Swagger for API documentation. The project is configured to use PostgreSQL as the database, and all configurations are customizable through environment variables.

## Prerequisites

Before getting started with the project, ensure that the following prerequisites are met:

1. Install the pre-commit package manager:

    ```bash
    sudo snap install pre-commit --classic
    ```

2. Install git hook scripts:

    ```bash
    pre-commit install
    ```

3. Install golangci-linter:

    ```bash
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
    ```

4. Install Swag Tool for API documentation:

    ```bash
    go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
    ```

## Environment Variables

Make sure to set the following environment variables according to your requirements:

- `SERVER_PORT`: Port on which the server will run (default: 8000).
- `BUILD_ENV`: Build environment (default: dev).

For database configuration:

- `DB_HOST`: PostgreSQL database host (default: localhost).
- `DB_PORT`: PostgreSQL database port (default: 5432).
- `DB_USERNAME`: PostgreSQL database username (default: postgres).
- `DB_PASSWORD`: PostgreSQL database password (default: 1234).
- `DB_NAME`: PostgreSQL database name (default: gin_service).

## Project Structure

The project structure follows the updated layout to improve organization and separation of concerns:

- `controllers/`: HTTP request handlers.
- `docs/`: Documentation files.
- `env/`: Environment variable management.
- `handlers/`: Logic for processing HTTP requests.
- `logger/`: Logging configuration and setup.
- `middlewares/`: Custom middleware functions.
- `models/`: Data structures and database models.
- `services/`: Database interactions logic.
- `utils/`: Utility functions.

## Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/gogintemplate.git
    cd gogintemplate
    ```

2. Update Swagger documentation:

    ```bash
    swag init
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

4. Access Swagger Documentation:

   Open your browser and navigate to:
    ```bash
    http://<SERVER_HOST>:<SERVER_PORT>/gin-service/api/v1/swagger/index.html
    ```

## Docker Support

A Dockerfile is provided to enable easy containerization of the application. Build and run the Docker container as follows:

```bash
docker build -t gogintemplate .
docker run -p <HOST_PORT>:<SERVER_PORT> gogintemplate
