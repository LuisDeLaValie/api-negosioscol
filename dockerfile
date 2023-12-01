# docker-compose up [--build]
FROM golang:1.21.4 AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/myapp

COPY go.mod .
RUN go mod download

COPY . .
EXPOSE 8080

ENTRYPOINT [ "go", "run", "main.go" ]

