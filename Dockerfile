# Start from the official Golang base image
FROM golang:1.22 as builder


# Set the working directory inside the container

WORKDIR /coursemanagement

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Copy the rest of the application's source code to the container

COPY . .


RUN go get
RUN go build -o /coursemanagement/build/myapp .
# Expose the port on which the application will run
EXPOSE 8080

ENTRYPOINT [ "/coursemanagement/build/myapp" ]
