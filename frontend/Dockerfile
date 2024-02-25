# Frontend build stage
FROM node:14 as frontend
WORKDIR /app/frontend

# Copy placeholder files (adjust as needed)
COPY placeholder/package.json .
COPY placeholder/public public

# Production-ready image
FROM nginx:alpine
COPY --from=frontend /app/frontend/public /usr/share/nginx/html