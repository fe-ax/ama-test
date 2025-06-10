# Time Server Demo

A simple Go HTTP server that responds with the current time.

## Features

- Returns current time in a readable format
- Health check endpoint
- Lightweight Docker container
- Request logging

## Endpoints

- `GET /` - Returns the current time
- `GET /health` - Health check endpoint

## Running Locally

### Prerequisites
- Go 1.21 or later

### Run the application
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Running with Docker

### Build the Docker image
```bash
docker build -t time-server .
```

### Run the container
```bash
docker run -p 8080:8080 time-server
```

## Running with Docker Compose

### Start the application
```bash
docker-compose up -d
```

### Stop the application
```bash
docker-compose down
```

## Deploying with Lagoon

This project is configured for deployment with [Lagoon](https://lagoon.readthedocs.io/).

### Prerequisites
- Access to a Lagoon instance
- `lagoon` CLI tool installed
- Project configured in Lagoon

### Deploy to Lagoon
```bash
# Add your project to Lagoon (one time setup)
lagoon add project

# Deploy the main branch
git push origin main
```

### Lagoon Features Configured
- **Basic Service**: Simple HTTP service deployment
- **Custom Routes**: Domain routing for your application

## Testing

Test the time endpoint:
```bash
curl http://localhost:8080
```

Test the health endpoint:
```bash
curl http://localhost:8080/health
```

## Example Response

```
Current time: 2024-01-15 14:30:45 UTC
``` 