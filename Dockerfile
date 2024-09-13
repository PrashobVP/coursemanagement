# Start from the official Golang base image
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /coursemanagement

# Copy the application's source code to the container
COPY . .

# Download dependencies and build the application
RUN go mod tidy
RUN go build -o /coursemanagement/build/myapp .

# Use a minimal base image for the final stage
FROM ubuntu:22.04

# Set the working directory
WORKDIR /coursemanagement

# Install MySQL client to test connectivity if needed (optional)
RUN apt-get update && \
    apt-get install -y mysql-client && \
    rm -rf /var/lib/apt/lists/*
    
    
# Install MySQL server
RUN apt-get update && \
DEBIAN_FRONTEND=noninteractive apt-get install -y mysql-server && \
rm -rf /var/lib/apt/lists/*



# Copy the built application from the builder stage
COPY --from=builder /coursemanagement/build/myapp /coursemanagement/build/myapp

# Expose the port on which the application will run
EXPOSE 8082


# Start MySQL and run the application
CMD /coursemanagement/build/myapp
