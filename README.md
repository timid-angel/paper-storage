# Getting Started

A simple storage server implemented using the native rpc library in Go and RabbitMQ for notifications using the publisher-subscriber pattern.

Before running the project, make sure to have an instance of RabbitMQ running and update its address in the environment variables file. To install and run RabbitMQ use:
```bash
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management
```

The project contains both the server and the client code. To start off, copy the template `.env.template`
```bash
cp .env.template .env
```

Then run the server using
```bash
make run
./bin/runner
```
Or simply
```
go run ./server
```

To run and interact with the sample client:
```bash
go run ./client
```

## Development

To facilitate development, a hot-reload tool for golang Fresh has been used. This will re-build and re-run the program when changes are made. The config settings for Fresh are in `/fresh.yaml`. To use fresh, simply run the command:
```bash
fresh
```
To run the service without using Fresh, use either of the options provided in the previous section.

# Project Structure
```
├───client
│       client.go
│       notification_handler.go
│       operation_handlers.go
│
├───config
│       env.go
│
├───docker
│       server.dockerfile
│       server.yml
│       wait-for-it.sh
│
├───domain
│   ├───dtos
│   │       dtos.go
│   │
│   └───entities
│           entities.go
│
└───server
    │   main.go
    │
    ├───controller
    │       controller.go
    │
    ├───domain
    │       domain.go
    │
    ├───repository
    │       storage_repository.go
    │
    ├───services
    │   └───rabbitmq
    │           rabbitmq.go
    │
    └───usecase
            paper_storage_usecase.go
```
> `client/` contains the client-side code for interacting with the key-value store.
 - `client.go`: Main client logic.
 - `notification_handler.go`: Handles notifications from RabbitMQ.
 - `operation_handlers.go`: Handles operations that are to be forwarded to the RPC server.

> `config/` contains configuration-related code.
 - `env.go`: Manages environment variables.

> `docker/` contains Docker-related files.
 - `server.dockerfile`: Dockerfile for the server.
 - `server.yml`: Docker Compose file for the server.
 - `wait-for-it.sh`: Script to wait for services to be ready.

> `domain/` contains the core business logic and domain models.
 - `dtos/`: Contains Data Transfer Objects.
    - `dtos.go`: DTO definitions.
 - `entities/`: Contains entity definitions.
    - `entities.go`: Entity definitions.

> `server/` contains the server-side code for the key-value store.
 - `main.go`: Entry point for the server application.
 - `controller/`: Manages the server controllers.
    - `controller.go`: Main controller logic for the RPC server.
 - `domain/`: Contains the core business logic and domain models.
    - `domain.go`: Main domain logic.
 - `repository/`: Contains the storage repository code.
    - `storage_repository.go`: Storage repository logic.
 - `services/`: Contains the service layer code.
    - `rabbitmq/`: Contains RabbitMQ service logic.
        - `rabbitmq.go`: RabbitMQ service logic.
 - `usecase/`: Contains the use case logic.
    - `paper_storage_usecase.go`: Use case logic for paper storage.

# Usage

The RPC server supports 4 major functions, which can be accessed using the following syntax from the client.

| Feature | Command |
| - | - | 
| Add Paper | `paperclient add <server-address> <author> <title> <path_to_file>` |
| List Papers | `paperclient list <server-address>` |
| Get Paper Meta Data | `paperclient detail <server-address> <number>` |
| Fetch Paper | `paperclient fetch <server-address> <number>` |