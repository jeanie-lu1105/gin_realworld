# syntax=docker/dockerfile:1
FROM golang:1.25 AS builder

WORKDIR /app
COPY . .
RUN go build .
EXPOSE 8080
CMD ./gin_realworld