# DamStudy (backend)

This is the backend for the DamStudy project. DamStudy is a study room finder for Oregon State University (OSU) students. The backend is written in Go and uses the MongoDB driver for Go, Websockets, and a RESTful API with the go-chi router.

## Prerequisites

- Air
- Go
- Docker
- Docker Compose
- MongoDB Atlas (if not using Docker)

## Getting Started

### MongoDB Atlas Setup

1. Create a free account on [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
2. Join the GDSC organization on MongoDB Atlas
3. Add the URI to the `internal/database/database.go` file
4. Run `make build` to build the application
5. Run `make run` to start the application
6. Navigate to `http://localhost:8080` in your browser

### Docker Setup (Recommended)

1. Clone the repository
2. Run `make docker-run` to start the MongoDB container
3. Run `make build` to build the application
4. Run `make run` to start the application
5. Navigate to `http://localhost:8080` in your browser

## Notes

- `init.js` is a script that will run when the MongoDB container is started. It will create the `damstudy` database and the `rooms` collection.
- Additionally, it "seeds" the database with some initial data.
- The `internal/database/database.go` file contains the connection to the MongoDB database. The URI is set as an environment variable.
- You can use the `make watch` command to live reload the application when changes are made.

> If you're still stuck, message @nyumat on the GDSC Discord server for help

## Makefile Commands

| Command             | Description                            |
| ------------------- | -------------------------------------- |
| `make all build`    | Run all make commands with clean tests |
| `make build`        | Build the application                  |
| `make run`          | Run the application                    |
| `make docker-run`   | Create DB container                    |
| `make docker-down`  | Shutdown DB container                  |
| `make watch`        | Live reload the application            |
| `make test`         | Run the test suite                     |
| `make clean`        | Clean up binary from the last build    |
| `make help`         | Show help message                      |
| `make docker-clean` | Clean up docker containers             |

## Deployment

(Coming soon)

## API Documentation

(Coming soon)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
