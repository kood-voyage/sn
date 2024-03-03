FROM golang:1.22

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o sn-api-go ./cmd/api

EXPOSE 6940

CMD ["./sn-api-go"]
