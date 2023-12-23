FROM golang:1.20

WORKDIR /app

COPY ./src/todoapp/go.mod ./src/todoapp/go.sum ./

RUN go mod download

COPY ./src/todoapp ./

EXPOSE 8080
