# Start from golang image
FROM golang:1.18 as base

RUN apt-get update && apt-get install -y ca-certificates openssl && update-ca-certificates
ADD ./proxy-golang-org2.pem /etc/ssl/certs/
RUN update-ca-certificates

ENV GO111MODULE=on
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum
WORKDIR /go/src/app
RUN go mod download