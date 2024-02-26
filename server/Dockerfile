# Backend build stage
FROM golang:1.17 as backend
WORKDIR /app/backend

# Copy placeholder files (adjust as needed)
COPY placeholder/main.go .

# Build Go application
RUN go build -o my-golang-app

# Production-ready image
FROM ubuntu:20.04
WORKDIR /app
COPY --from=backend /app/backend/my-golang-app /app/my-golang-app

EXPOSE 8080
CMD ["./my-golang-app"]