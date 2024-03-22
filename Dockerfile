# Use a base image with necessary runtime environments
FROM ubuntu:latest

RUN apt-get update
# # Install Java
# RUN apt-get install -y default-jdk

# # Install Python
# RUN apt-get install -y python3

# # Install Golang
# RUN apt-get install -y golang-1.20

# # Set Go environment variables
# ENV PATH="/usr/lib/go-1.17/bin:${PATH}"
# ENV GOPATH="/go"
# ENV PATH="${GOPATH}/bin:${PATH}"

# # Install Node.js and npm for JavaScript
# RUN apt-get install -y nodejs npm

# # Install C++ compiler
# RUN apt-get install -y build-essential

# # Set working directory
# WORKDIR /app

# COPY go.mod go.sum ./

# # Download and install dependencies
# RUN go mod download

# # Copy your code files into the container
# COPY ./ ./

# # Assigning environment variables
# ARG PORT 8080
# ARG DB_URI

# ENV PORT=8080\
#     DB_URI="mongodb://localhost:27017"

# # Build the Go application
# RUN go build -o server.exe ./app

# EXPOSE 8080

# # Command to run the executable
# CMD ["./server"]

RUN apt-get install -y golang
RUN go version