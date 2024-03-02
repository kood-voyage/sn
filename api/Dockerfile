# build Stage
FROM golang:alpine AS build

# installing build dependencies 
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY . .

WORKDIR /app/cmd/api

RUN CGO_ENABLED=1 go build -o sn-go-api .

# finaal Stage
FROM alpine:latest

RUN mkdir /app

COPY --from=build /app/cmd/api/sn-go-api /app/

WORKDIR /app

COPY . /app/

WORKDIR /app

COPY config/ config/

EXPOSE 8080

CMD ["./sn-go-api"]
