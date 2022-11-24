# TODO: hot reload
FROM golang:1.19

WORKDIR /app

RUN apt-get update && apt-get install sqlite3 && go install github.com/swaggo/swag/cmd/swag@latest

CMD ["tail", "-f", "/dev/null"]