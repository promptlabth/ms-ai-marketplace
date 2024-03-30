FROM golang:1.22 as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/*.go


FROM alpine:3.17.0

# Set work directory
WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata libc6-compat

# Set timezone
ENV TZ=Asia/Bangkok

# Copy build result from builder
COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
