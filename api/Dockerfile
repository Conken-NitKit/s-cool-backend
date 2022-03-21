FROM golang:1.17 

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o bin/app ./cmd/main.go
