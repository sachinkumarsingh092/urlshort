# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.15.4 as builder

# Ensure /app directory exists.
RUN mkdir -p /app

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o urlshort

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates bash

# Expose the port 8080 to be used for port forwarding.
EXPOSE 8080

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/urlshort /urlshort

# Copy the urlmaps yaml files in /urlmaps directory in alpine.
COPY --from=builder /app/urlmaps ./urlmaps

# Run the web service on container startup.
ENTRYPOINT [ "/urlshort" ]