# # Use a base image that includes the C compiler and other build tools
# FROM golang:1.22-alpine as builder

# # Install GCC and other necessary tools
# RUN apk add --no-cache gcc musl-dev

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the Go Modules manifests
# COPY go.mod go.sum ./
# # Download Go module dependencies
# RUN go mod download

# # Copy the source code into the container
# COPY . .

# # Set environment variable to enable CGO
# ENV CGO_ENABLED=1

# # Build the Go app
# RUN go build -o /main ./cmd/

# # Start a new stage from scratch
# FROM alpine:latest  
# WORKDIR /

# COPY prompt-lab-383408-512938be4baf.json ./prompt-lab-383408-512938be4baf.json
# COPY --from=builder /main ./
# # COPY .env ./.env
# EXPOSE 8080
# # Command to run the executable
# CMD ["/main"]




FROM golang:1.22 as builder

RUN apk --no-cache add ca-certificates tzdata libc6-compat

WORKDIR /app

COPY . ./

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/*.go


FROM alpine:3.17.0

# Set work directory
WORKDIR /app



# Set timezone
ENV TZ=Asia/Bangkok

COPY prompt-lab-383408-512938be4baf.json ./prompt-lab-383408-512938be4baf.json

COPY --from=builder /main ./
# COPY .env ./.env
EXPOSE 8080

ENTRYPOINT ["./main"]
