FROM golang:1.19.0

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy
