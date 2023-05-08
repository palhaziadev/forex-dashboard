# Start from golang image
FROM golang:1.18 as base

RUN apt-get update && apt-get install -y ca-certificates openssl && update-ca-certificates
ADD ./proxy-golang-org7.pem /etc/ssl/certs/
ADD ./sum-golang-org2.pem /etc/ssl/certs/

ADD ./registry-1-docker-io1.pem /etc/ssl/certs/

ADD ./proxy-golang-org5.pem /etc/ssl/certs/
ADD ./proxy-golang-org6.pem /etc/ssl/certs/
ADD ./proxy-golang-org7.pem /etc/ssl/certs/
ADD ./sum-golang-org.pem /etc/ssl/certs/
ADD ./sum-golang-org1.pem /etc/ssl/certs/
ADD ./sum-golang-org2.pem /etc/ssl/certs/
RUN update-ca-certificates

ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum
WORKDIR /go/src/app
RUN go mod download