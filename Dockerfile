# Use a base image that includes the Go compiler and other build tools
FROM golang:1.23 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Start a new stage from scratch
FROM alpine:3.17.0

# Set work directory
WORKDIR /app

# Install necessary packages
RUN apk --no-cache add ca-certificates tzdata libc6-compat

# Set timezone
ENV TZ=Asia/Bangkok

# Copy the necessary files from the builder stage
COPY --from=builder /app/main ./
COPY prompt-lab-cred.json ./
COPY firebase-credential.json ./
# Expose port 8080
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./main"]
