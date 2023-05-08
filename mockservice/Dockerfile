# Start from golang base image
FROM base2 as mockservice

ENV PROJECT_DIR=/go/src/app/mockservice

WORKDIR /go/src/app/mockservice

COPY ./pkg /go/src/app/pkg
COPY ./mockservice /go/src/app/mockservice

ENTRYPOINT CompileDaemon -build="go build -o /go/bin/mockservice" -command="/go/bin/mockservice"