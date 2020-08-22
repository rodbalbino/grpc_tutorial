
#build stage
FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update add --no-cache \
    git bash ca-certificates gcc g++ libc-dev

RUN mkdir /server
RUN mkdir -p /server/proto

WORKDIR /server

COPY ./proto/service.pb.go /server/proto
COPY ./server/main.go /server

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o server .

ENTRYPOINT ./server
LABEL Name=grpc-calc Version=0.0.1
EXPOSE 4040
