# Start from the official Golang base image
FROM golang:1.22 as builder

# Set environment variables 
# ENV DB_HOST=host.docker.internal 
# ENV DB_PORT=5432 
ENV DB_USER=root
ENV DB_PASSWORD=1234 
# ENV DB_NAME=mydatabase
# Set the working directory inside the container

WORKDIR /coursemanagement

# Copy the rest of the application's source code to the container

COPY . .


RUN go get
RUN go mod tidy
RUN go build -o /coursemanagement/build/myapp .
# Expose the port on which the application will run
EXPOSE 8080

ENTRYPOINT [ "/coursemanagement/build/myapp" ]
