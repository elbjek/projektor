# Start from golang v1.11 base image
FROM golang:1.11

# Build Args
ARG APP_NAME=hakaton
ARG LOG_DIR=/${APP_NAME}/logs

# Create Log Directory
RUN mkdir -p ${LOG_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/callicoder/go-docker

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Declare volumes to mount
VOLUME ["/hakaton/logs"]

# Run the binary program produced by `go install`
CMD ["./hakaton"]
