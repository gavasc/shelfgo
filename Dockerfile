# Build Stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o shelfgo .

# Final Image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/shelfgo .
CMD ["./shelfgo"]