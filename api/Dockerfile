FROM golang:1.17 

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /build
COPY . /build

EXPOSE 80
RUN go mod tidy && go build 
CMD ["go", "run", "main.go"]
