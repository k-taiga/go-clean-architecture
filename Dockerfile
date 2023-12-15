FROM golang:1.20

WORKDIR /app

COPY ./src/go-testing/go.mod ./src/go-testing/go.sum ./

RUN go mod download

COPY ./src/go-testing ./

EXPOSE 8080
