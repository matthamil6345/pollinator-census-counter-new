# Use the official Golang image to create a build artifact.
# This is the base image for the first stage of the multi-stage build.
FROM golang:1.22.4 as builder
WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o /bin/server ./cmd/server
EXPOSE 8080
ENTRYPOINT [ "/bin/server" ]

# Set the Current Working Directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # Copy the source from the current directory to the Working Directory inside the container
# COPY . .

# # Build the Go app
# RUN go build -o . main.go

# # Expose port 8080 to the outside world
# EXPOSE 8080

# # Command to run the executable
# CMD ["./main"]
