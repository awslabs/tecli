FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go get ./...
RUN go generate ./...

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o tecli .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/tecli .

# Cleanup
RUN rm -rf /build

# Command to run when starting the container
CMD ["/dist/tecli"]