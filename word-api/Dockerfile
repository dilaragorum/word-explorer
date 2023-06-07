FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download

RUN CG0_ENABLED=0 go build -o word cmd/api/main.go

FROM alpine:latest

COPY --from=builder /app/word .

EXPOSE 3200

CMD [ "./word" ]