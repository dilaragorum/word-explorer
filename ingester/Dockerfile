FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download

RUN CG0_ENABLED=0 go build -o ingester cmd/main.go


FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ=Europe/Istanbul

COPY --from=builder /app/ingester .
COPY --from=builder /app/.env .
COPY --from=builder /app/credentials.json .
COPY --from=builder /app/token.json .

EXPOSE 3100

CMD [ "./ingester" ]