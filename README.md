# maker-checker

## Overview

This is a simple REST API implementation of a maker-checker pattern. The idea is to have a maker create a request and a checker approve or reject it.
The maker and checker can be the same person or different people.

## Getting Started

### Prerequisites

- Install Go 1.23
- Ensure docker is installed and running on your machine
- Check if `make` command is available by running `make --version`. If not, use either methods based on your OS: 
```shell
# For Ubuntu/Debian based linux
sudo apt install build-essential

# For Mac
xcode-select --install
```
### Build and Run
```shell
# Download go dependencies
make go-deps
# Build the project 
make build
# Start db instance in docker
make services-up
# Start server on port 8080
make http-start
```

## HTTP endpoints

| Endpoint                   | Method | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Description                                                             |                                                                        
|----------------------------|--------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| `/health`                  | GET    | curl -L 'http://localhost:8080/health'                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Health check for service. Responds with `{"status": "ok"}` in the body. |                   
| `/profiles/`               | POST   | curl -L 'http://localhost:8080/profiles/' -H 'Content-Type: application/json' --data-raw '{"firstName":"Su","lastName":"Shi","email":"sushi@foodiecrush.com"}'<br/><br/>curl -L 'http://localhost:8080/profiles/' -H 'Content-Type: application/json' --data-raw '{"firstName":"Kim","lastName":"Chi","email":"kimchi@foodiecrush.com"}'<br/><br/>curl -L 'http://localhost:8080/profiles/' -H 'Content-Type: application/json' --data-raw '{"firstName":"Nasi","lastName":"Lemak","email":"nasi.lemak@foodiecrush.com"}' | Create new user profile. Returns `{"profile_id": "<uuid>"}`.            | 
| `/profiles/{id}`           | GET    | curl -L 'http://localhost:8080/profiles/<uuid>'                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Get profile by ID.                                                      | 
| `/documents/`              | POST   | curl -L 'http://localhost:8080/documents/' -H 'Content-Type: application/json' -d '{"creatorID":"<profileID>","description":"Test document","status":"Draft"}'                                                                                                                                                                                                                                                                                                                                                            | Create new document.                                                    |
| `/documents/{id}`          | GET    | curl -L 'http://localhost:8080/documents/<documentID>'                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Get document by ID.                                                     |
| `/documents/requests/`     | POST   | curl -L 'http://localhost:8080/documents/requests/' -H 'Content-Type: application/json' -d '{"documentID":"<documentID>","creatorID":"<user1ProfileID>","approvers":[{"id":"<user2profileID>"},{"id":"<user3profileID>"}],"recipientEmail":"recipient@example.com"}'                                                                                                                                                                                                                                                      | Create new document request.                                            |
| `/documents/requests/{id}` | PUT    | curl -L 'http://localhost:8080/documents/requests/' -H 'Content-Type: application/json' -d '{"documentID":"<documentID>","creatorID":"<user1ProfileID>","approvers":[{"id":"<user2profileID>","status":"Approved"},{"id":"<user3profileID>","status":"Approved"}],"recipientEmail":"recipient@example.com"}'                                                                                                                                                                                                              | Update document request by ID. Use to approve/reject the request.       |
| `/documents/requests/{id}` | GET    | curl -L 'http://localhost:8080/documents/requests/<docRequestID>'                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Get document request by ID.                                             |

## Creating migrations

Use the migrate tool in `/tools` directory or follow the instructions on [this page](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation) to install on your system.

In order to create migrations, the following can be run:

```bash
$ make migration-create NAME=your_migration_name
```
This should create 2 new files in the `store/psql/migrations` directory with a timestamp, name, and their up/down variants.

## Running Tests
```shell
# Install gotestsum
go install gotest.tools/gotestsum@latest
# Run tests  
make go-unit-test
```

## Project Structure
```shell
.
├── build
│   └── maker-checker
├── cmd      # Main application
│   └── maker-checker
│       ├── config
│       │   └── config.go
│       └── main.go
├── logger    # Custom logger that includes context information
│   ├── contextlogger.go
│   ├── logger.go
│   └── zap.go
├── pkg       # Core application logic
│   ├── document
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── request
│   │   │   ├── model.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── service.go
│   ├── helpers.go
│   └── profile
│       ├── model.go
│       ├── repository.go
│       └── service.go
├── store       # Database connection setup and repository implementations
│   └── psql      # Implies the storage technology used i.e. Postgresql
│       ├── db.go    # Connection and migration setup
│       ├── document
│       │   ├── repository.go
│       │   └── request
│       │       └── repository.go
│       ├── init   # Contains database setup scripts
│       │   └── create_database.sql
│       ├── migrations
│       │   ├── 20241230052529_create_table_user.down.sql
│       │   ├── 20241230052529_create_table_user.up.sql
│       │   ├── 20250101082609_create_table_document.down.sql
│       │   ├── 20250101082609_create_table_document.up.sql
│       │   ├── 20250102152249_create_table_document_request.down.sql
│       │   └── 20250102152249_create_table_document_request.up.sql
│       └── profile
│           └── repository.go
└── transport   # Contains all transport related code. Currently only HTTP but can be expanded to include gRPC, graphql, rabbitmq/AMQP etc.
    └── http
        ├── api.go
        ├── decode
        │   └── decode.go       # Decodes request body into a struct
        ├── endpoints.go        # Delegates request to corresponding service under `/pkg`
        ├── middleware.go       # Contains middleware functions. Currently only logging middleware.
        ├── router.go           # Contains all routes and their corresponding handlers
        └── types
            └── types.go

```

## Libraries used
- Go-kit: https://github.com/go-kit/kit
- Viper: https://github.com/spf13/viper
- Zap: https://github.com/uber-go/zap
- Sqlx: https://github.com/jmoiron/sqlx
- Migrate: https://github.com/golang-migrate/migrate


## TODO
- Validate request body
- Implement authentication and authorization
- Implement instrumentation
- Add more tests
