
# Word Explorer

Main dependencies:
- HTTP Server: [fiber](https://github.com/gofiber/fiber)
- Database : [elasticsearch](https://github.com/elastic/go-elasticsearch)
- Message Broker: [kafka](https://github.com/segmentio/kafka-go)


## Project Layout

```
.
├── README.md
├── consumer
│   ├── Dockerfile
│   ├── cmd
│   │   └── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   └── vocabulary
│   │       ├── consumer.go
│   │       ├── model.go
│   │       └── service.go
│   └── pkg
│       └── kafka_reader.go
├── docker-compose.yml
├── ingester
│   ├── Dockerfile
│   ├── cmd
│   │   └── main.go
│   ├── credentials.json
│   ├── go.mod
│   ├── go.sum
│   ├── hash.txt
│   ├── internal
│   │   └── vocabulary
│   │       ├── google-sheet-client.go
│   │       ├── model.go
│   │       ├── publisher.go
│   │       └── service.go
│   ├── pkg
│   │   ├── cron_client.go
│   │   ├── hash.go
│   │   ├── kafka_writer.go
│   │   └── string.go
│   └── token.json
└── word-api
    ├── Dockerfile
    ├── cmd
    │   └── api
    │       └── main.go
    ├── go.mod
    ├── go.sum
    ├── internal
    │   └── vocabulary
    │       ├── handler.go
    │       ├── model.go
    │       ├── repository.go
    │       ├── request.go
    │       ├── response.go
    │       └── service.go
    ├── pkg
    │   └── elastic
    │       └── elastic.go
    └── version1-2023-06-01-mapping.json


17 directories, 37 files
```
## Project Architecture

![Screenshot](Project-Architecture.png)



