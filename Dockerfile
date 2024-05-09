# First stage: build the Go binary.
FROM golang:latest as builder

# Set the working directory inside the builder container.
WORKDIR /app

# Copy the source code into the builder container.
COPY . .

# Initialize Go modules and fetch dependencies.
RUN go mod init app
RUN go mod tidy

# Build the Go application to a binary named 'app'.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Second stage: create a lightweight deployment image.
FROM alpine:latest  

# Add ca-certificates in case you need HTTPS.
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder stage.
COPY --from=builder /app/app .

# Command to run the binary.
CMD ["./app"]

