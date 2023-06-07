
# Word Explorer

Main dependencies:
- HTTP Server: [fiber](https://github.com/gofiber/fiber)
- Database : [elasticsearch](https://github.com/elastic/go-elasticsearch)
- Message Broker: [kafka](https://github.com/segmentio/kafka-go)
- Cron: [gocron](https://github.com/go-co-op/gocron)


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

## Motivation

Imagine you have english words with their meanings and example sentences in your excel sheet in your drive. However, because you are really hardworking, these sheet's size is expanding day by day. You can not search word quickly because sheet is frozen. As a result, you need application to search your words quickly.

Our aim is to transform data from excel sheet to elastic search to make fast search. Therefore, let's know our applications:

Ingester: Retrieve data from excel sheet and compare them. If there is any change, it sends changed data as message to Kafka broker.

Consumer: Read messages from Kafka Broker and sends these messages to Word API

Word API: Enable us to communicate Elastic Search which provide full text search very fast and store our data.

All these applications are dockerized and docker-composed file is used to run Elastic Search, Kafka, Kibana, Zookeper containers.
