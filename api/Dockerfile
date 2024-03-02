# Build Stage
FROM golang:alpine AS build

# Installing build dependencies if needed
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY . .

WORKDIR /app/cmd/api

RUN CGO_ENABLED=1 go build -o sn-go-api .

# Final Stage
FROM alpine:latest

RUN mkdir /app

COPY --from=build /app/cmd/api/sn-go-api /app/

WORKDIR /app

# Copy the necessary config files
COPY config/ config/

EXPOSE 8080

CMD ["./sn-go-api"]
