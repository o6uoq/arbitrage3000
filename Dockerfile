# Use the official Golang image as base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Remove any previously initialized go.mod and go.sum files
# (this is in case the container data wasn't destroyed)
RUN rm -f go.mod && rm -f go.sum

# initialize Go modules
RUN go mod init app

# fetch dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o app .

# Run the compiled binary when the container starts
CMD ["./app"]

