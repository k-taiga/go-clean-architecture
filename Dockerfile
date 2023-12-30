FROM golang:1.20

WORKDIR /app

COPY ./src/sluck/go.mod ./src/sluck/go.sum ./

RUN go mod download

COPY ./src/sluck ./

EXPOSE 8080
