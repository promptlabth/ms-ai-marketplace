# ms-ai-marketplace

# run windows use below dockerfile

```docker
FROM golang:1.22-alpine as builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1

RUN go build -o /main ./cmd/

FROM alpine:latest  

WORKDIR /

COPY --from=builder /main ./

CMD ["/main"]
```