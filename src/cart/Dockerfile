FROM golang:1.21-alpine3.18 AS builder

RUN apk add --no-cache git

# Set the Working Directory inside the container.
# WORKDIR /tmp/src/app
RUN mkdir -p /tmp/src/app
WORKDIR /tmp/src/app

# Copy root module with generated proto files.
COPY go.mod /tmp
COPY gen /tmp/gen

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and go.sum.
COPY src/cart/go.* /tmp/src/app/
RUN go mod download

# Copy local code to the container image.
COPY src/cart/ .

# Build the binary.
RUN go build -o server .

# Start fresh from a smaller image.
FROM alpine:3.18
RUN apk add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /tmp/src/app/server /app/server

# Run the service on container startup.
CMD ["/app/server"]
