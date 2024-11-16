# Build Stage
FROM golang:1.20-alpine AS builder
WORKDIR /app

# Install required tools
RUN apk add --no-cache git

# Copy application files and dependencies
COPY . .
RUN go mod tidy
RUN go build -o myproject .

# Final Stage
FROM alpine:latest
WORKDIR /root/

# Copy the compiled application and store_master.csv
COPY --from=builder /app/myproject .
COPY store_master.csv .

# Expose the port for the application
EXPOSE 8080

# Command to run the application
CMD ["./myproject"]
