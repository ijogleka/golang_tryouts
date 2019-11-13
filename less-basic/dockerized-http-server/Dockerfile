# Get a image from docker hub for setting up a basic golang environment
FROM golang:1.8-onbuild

# Who will be maintaining this container?
MAINTAINER ijogleka@gmail.com

# Working directory inside the container
WORKDIR /app

# Copy over the package configuration file
COPY go.mod go.sum ./

# Download all the required packages (with specific versions) as defined in go.mod
RUN go mod download

# Copy over the source files for my app!
COPY . .

# Build a binary from the source files (and the downloaded package dependencies)
# inside the container
# Call the binary executable, `main`
RUN go build -o main .

# Start the container by running the generated binary
CMD ["./main"]