# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Ayesha Kaleem <akaleem306@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /web_server

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build main.go
RUN chmod +x ./main.go

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable

CMD ["./main"]
