# Grab golang package
FROM golang:latest

# Set the Current Working Directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go app
RUN go build main.go

# Expose port 8081
EXPOSE 8081

# Run the executable
CMD ["./main"]
