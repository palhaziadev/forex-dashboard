# Start from golang base image
FROM base as server

WORKDIR /go/src/app

COPY ./pkg /go/src/app/pkg
COPY ./server /go/src/app/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/server /go/src/app/server

ENTRYPOINT [ "/go/bin/server" ]


