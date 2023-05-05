FROM golang:1.20-alpine3.17

WORKDIR /app
COPY . /app

# go get で git を使うため追加
RUN apk upgrade --update && \
  apk --no-cache add git && \
  apk add curl

RUN go mod download

# install air, sqlc
RUN go install github.com/cosmtrek/air@latest && \
  go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
  mv /app/migrate.linux-amd64 /usr/bin/migrate && \
  which migrate

EXPOSE 8080
CMD [ "air", "-c", ".air.toml" ]
