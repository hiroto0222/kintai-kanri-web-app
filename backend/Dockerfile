FROM golang:1.20-alpine3.17

WORKDIR /app
COPY . /app

# go get で git を使うため追加
RUN apk upgrade --update && \
  apk --no-cache add git

RUN go mod download

# install air
RUN go install github.com/cosmtrek/air@latest && \
  go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

EXPOSE 8080
CMD [ "air", "-c", ".air.toml" ]
