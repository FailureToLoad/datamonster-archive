FROM golang:1.22.6-alpine3.20 AS build-stage

WORKDIR /app

# Copy the source code and download dependencies
COPY ./ ./
RUN go mod download

# Build the application binary
RUN go build -o /out/apiserver ./cmd/apiserver/main.go
RUN ls /out


FROM gcr.io/distroless/base-debian12:latest AS run-stage

WORKDIR /
COPY --from=build-stage /out/apiserver /apiserver
USER nonroot:nonroot

EXPOSE 8080

ARG AZURE_CLIENT_ID
ARG AZURE_TENANT_ID
ARG AZURE_CLIENT_SECRET
ARG VAULT_URI

# Set environment variables
ENV MODE=cluster
ENV AZURE_CLIENT_ID=${AZURE_CLIENT_ID}
ENV AZURE_TENANT_ID=${AZURE_TENANT_ID}
ENV AZURE_CLIENT_SECRET=${AZURE_CLIENT_SECRET}
ENV VAULT_URI=${VAULT_URI}

ENTRYPOINT ["./apiserver"]