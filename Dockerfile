# builder image
FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o producer ./cmd/producer
RUN go build -o consumer ./cmd/consumer

# run image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/run_web_server .

EXPOSE 8080

# never separate the command and output redirection
CMD [ "/bin/sh", "-c", "/app/consumer > /var/log/gateway.log" ]


