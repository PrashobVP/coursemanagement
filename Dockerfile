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
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /coursemanagement

# Install MySQL
RUN apt-get update && \
    apt-get install -y default-mysql-server && \
    rm -rf /var/lib/apt/lists/*

# Set environment variables
#ENV MYSQL_ROOT_PASSWORD=1234
#ENV MYSQL_DATABASE=mydatabase
ENV DB_USER=root
ENV DB_PASSWORD=1234
#ENV DB_HOST=localhost
#ENV DB_PORT=3306
#ENV DB_NAME=mydatabase

# Copy the built application from the builder stage
COPY --from=builder /coursemanagement/build/myapp /coursemanagement/build/myapp

# Expose the port on which the application will run
EXPOSE 8080
EXPOSE 3306

# Start MySQL and run the application
CMD service mysql start && \
    mysql -u root -p1234 -e "CREATE DATABASE IF NOT EXISTS coursemanagement;" && \
    /coursemanagement/build/myapp
