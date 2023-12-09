# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -v -o main ./cmd

# Set environment variables
ENV DEV=true
ENV DEBUG=true
ENV DATABASE_NAME=real_estate
ENV DATABASE_PASSWORD_HASH=test
ENV ELASTIC_SEARCH_URL=http://localhost:9001/api/v1/jobs
ENV JWT_SECRET=test
ENV SERVER_PORT=8001
ENV SERVER_TIMEOUT=30
ENV ORGANIZATION_EMAIL_EMAIL=nahomasnakeaddis@gmail.com
ENV ORGANIZATION_EMAIL_PASSWORD="wihm nxhf lehl jzsp"
ENV DB_HOST=db
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV PORT=5432

# Expose the server port
EXPOSE 8001

# Run the Go application
CMD ["./main"]