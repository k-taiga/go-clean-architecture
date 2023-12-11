FROM golang:1.20

WORKDIR /app

COPY ./src/1-editing /app/

EXPOSE 8080