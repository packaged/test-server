# syntax=docker/dockerfile:1
FROM golang:1.19-alpine3.16 AS builder

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server

FROM scratch

COPY --from=builder /app/server /app/server
WORKDIR /app

ENTRYPOINT [ "./server" ]
