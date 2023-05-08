# Start from golang base image
# FROM base as server

# WORKDIR /go/src/app

# COPY ./pkg /go/src/app/pkg
# COPY ./server /go/src/app/server

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/server /go/src/app/server

# ENTRYPOINT [ "/go/bin/server" ]
FROM base2 as server

ENV PROJECT_DIR=/go/src/app/server

WORKDIR /go/src/app/server

COPY ./pkg /go/src/app/pkg
COPY ./server /go/src/app/server

ENTRYPOINT CompileDaemon -build="go build -o /go/bin/server" -command="/go/bin/server" -graceful-kill=true
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/server /go/src/app/server

# ENTRYPOINT [ "/go/bin/server" ]



