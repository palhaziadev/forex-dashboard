# # Start from golang base image
# FROM base as dbservice

# WORKDIR /go/src/app

# COPY ./pkg /go/src/app/pkg
# COPY ./dbservice /go/src/app/dbservice

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/dbservice /go/src/app/dbservice

# ENTRYPOINT [ "/go/bin/dbservice" ]

# Start from golang base image
FROM base2 as dbservice

ENV PROJECT_DIR=/go/src/app/dbservice

WORKDIR /go/src/app/dbservice

COPY ./pkg /go/src/app/pkg
COPY ./dbservice /go/src/app/dbservice

ENTRYPOINT CompileDaemon -build="go build -o /go/bin/dbservice" -command="/go/bin/dbservice"


