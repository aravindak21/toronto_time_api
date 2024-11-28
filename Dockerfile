# Build stage
FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Run stage
FROM ubuntu:20.04
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
