# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="DUNA"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN apk add --no-cache bash curl

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download


# Copy the source from the current directory to the working Directory inside the container
COPY . .

ENV GIN_MODE="release"

EXPOSE 80

# Build the Go app
#RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#
#
## Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash curl

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .

ENV GIN_MODE="release"
ENV GIN_PORT=":80"

# Expose port 8080 to the outside world
EXPOSE 80

#Command to run the executable
ENTRYPOINT ["./main"]