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
RUN go build main.go


FROM scratch
COPY --from=build /go/src/myapp /go/bin/myapp
ENTRYPOINT ["/go/bin/myapp"]