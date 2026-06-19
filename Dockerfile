FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.cn,direct
RUN go build .
EXPOSE 8080
CMD ./gin_realworld