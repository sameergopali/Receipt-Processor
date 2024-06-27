# Use an official Go runtime as a parent image
FROM golang:1.21.1  as builder

# Ensure Go compiler builds a static linked binary
ENV CGO_ENABLED=0
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod 
COPY go.mod  ./
RUN go mod download

# Install make and git
RUN apt update && apt install  make git

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN make build

# Staging 
FROM alpine:latest 

RUN mkdir /config
COPY --from=builder /app/main /usr/local/bin/
COPY --from=builder /app/config/config.yml /config/config.yml
# Give Execution permission
RUN chmod +x /usr/local/bin/main

# Expose port 808
EXPOSE  8080

# Entrypoint
CMD ["main"]
