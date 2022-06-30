# Start from golang base image
FROM base as mockservice

WORKDIR /go/src/app

COPY ./pkg /go/src/app/pkg
COPY ./mockservice /go/src/app/mockservice

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/mockservice /go/src/app/mockservice

ENTRYPOINT [ "/go/bin/mockservice" ]