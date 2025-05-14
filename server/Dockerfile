# Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# Runtime
FROM alpine:3.21

# Install timezone data
RUN apk add --no-cache tzdata curl

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3000

CMD ["./server"]