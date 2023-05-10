# Build Stage
FROM golang:1.20-alpine3.17 AS builder

WORKDIR /app
COPY . .

# build executable go
RUN go build -o main main.go
# go get で git を使うため追加
RUN apk upgrade --update && \
  apk --no-cache add git && \
  apk add curl
# install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY .local.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./migrations

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
