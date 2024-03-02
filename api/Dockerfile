#Build
FROM golang:alpine as build

# installing build dependencies
RUN apk add --no-cache gcc musl-dev


RUN mkdir /app

WORKDIR /app

COPY . .

WORKDIR /app/cmd/api

RUN CGO_ENABLED=1 go build -o sn-go-api .

#Image

FROM alpine:latest

RUN mkdir /db

RUN mkdir /app

COPY --from=build /app /app

WORKDIR /app/cmd/api

EXPOSE 8080:8080

CMD ./sn-go-api
