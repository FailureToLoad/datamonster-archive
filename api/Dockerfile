FROM golang:1.23.2-alpine3.20 AS build-stage

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

ENV MODE=prod

ENTRYPOINT ["./apiserver"]