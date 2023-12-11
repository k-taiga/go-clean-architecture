FROM golang:1.20

WORKDIR /app

COPY ./src/1-editing/go.mod ./src/1-editing/go.sum ./

RUN go mod download

COPY ./src/1-editing ./

EXPOSE 8080
