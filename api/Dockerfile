FROM golang:1.22.6-alpine3.20 AS build-stage

WORKDIR /app

# Copy the source code and download dependencies
COPY ./ ./
RUN go mod download

# Build the application binary
RUN go build -o /out/apiserver ./cmd/apiserver/main.go
RUN ls /out


# Stage 2: Set up the runtime environment using distroless
FROM alpine:latest AS run-stage

# Set the working directory
WORKDIR /

RUN apk add --no-cache openssl && \
    wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-alpine-linux-amd64-v0.6.1.tar.gz && \
    tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-v0.6.1.tar.gz && \
    rm dockerize-alpine-linux-amd64-v0.6.1.tar.gz


COPY --from=build-stage /out/apiserver /apiserver

# Non-root user setup for security
RUN adduser -D nonroot && \
    chown nonroot /apiserver

USER nonroot
# Expose the application port
EXPOSE 8080

# Set environment variables
ENV MODE=cluster
ENV AZURE_CLIENT_ID=""
ENV AZURE_TENANT_ID=""
ENV AZURE_CLIENT_SECRET=""
ENV VAULT_URI=""

ENTRYPOINT ["dockerize", "-wait", "tcp://dm-records:5432", "-timeout", "30s", "/apiserver"]