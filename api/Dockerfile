FROM golang:1.22

RUN apt-get update && \
    apt-get install -y sqlite3 libsqlite3-dev

WORKDIR /go/src/app/cmd/api

COPY . .

RUN go mod download

RUN go build -o sn-api-go .

EXPOSE 6940

CMD ["./sn-api-go"]
