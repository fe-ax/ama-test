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