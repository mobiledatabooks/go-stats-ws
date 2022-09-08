# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Copy local code to the container image.
COPY . ./
COPY go.mod ./

# Build the binary.
RUN go build -v main.go
RUN go test -v -cover ./...

FROM gcr.io/distroless/base
########################

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/main /app/main

# Run the web service on container startup.
CMD ["/app/main"]