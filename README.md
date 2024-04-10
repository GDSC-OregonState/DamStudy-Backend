# DamStudy (backend)

This is the backend for the DamStudy project. DamStudy is a study room finder for Oregon State University (OSU) students. The backend is written in Go and uses the MongoDB driver for Go, Websockets, and a RESTful API with the go-chi router.

## Prerequisites

- Air
- Go
- Docker
- Docker Compose
- MongoDB (if not using Docker)

## Getting Started

1. Clone the repository
2. Run `make build` to build the application
3. Run `make docker-run` to start the MongoDB container
4. Run `make run` to start the application
5. Navigate to `http://localhost:8080` in your browser
6. Run `make docker-down` to stop the MongoDB container

## Makefile Commands

| Command            | Description                            |
| ------------------ | -------------------------------------- |
| `make all build`   | Run all make commands with clean tests |
| `make build`       | Build the application                  |
| `make run`         | Run the application                    |
| `make docker-run`  | Create DB container                    |
| `make docker-down` | Shutdown DB container                  |
| `make watch`       | Live reload the application            |
| `make test`        | Run the test suite                     |
| `make clean`       | Clean up binary from the last build    |

## Deployment

(Coming soon)

## API Documentation

(Coming soon)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
