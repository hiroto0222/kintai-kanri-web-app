FROM golang:1.20-alpine3.17

WORKDIR /app
COPY . /app

# add uuid to postgresql
ENV POSTGRES_EXTENSIONS=uuid-ossp

# go get で git を使うため追加
RUN apk upgrade --update && \
  apk --no-cache add git && \
  apk add curl

RUN go mod download

# install air
RUN go install github.com/cosmtrek/air@latest

# install sqc (error with go get: the PostgreSQL engine requires cgo. Please set CGO_ENABLED=1)
RUN curl -L https://github.com/kyleconroy/sqlc/releases/download/v1.18.0/sqlc_1.18.0_linux_amd64.tar.gz | tar xvz && \
  mv /app/sqlc /usr/bin/sqlc && \
  which sqlc

# install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
  mv /app/migrate.linux-amd64 /usr/bin/migrate && \
  which migrate

EXPOSE 8080
CMD [ "air", "-c", ".air.toml" ]
