FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download

RUN CG0_ENABLED=0 go build -o consumer cmd/main.go

FROM alpine:latest

COPY --from=builder /app/consumer .

EXPOSE 3000

CMD [ "./consumer" ]