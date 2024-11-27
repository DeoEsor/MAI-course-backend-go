FROM golang:1.22.4

# Set destination for COPY
WORKDIR /app

RUN go install github.com/go-task/task/v3/cmd/task@latest

# Download Go modules
COPY go.mod go.sum ./
COPY Taskfile.yml .
RUN go mod download

COPY . ./

# Build
RUN task build

EXPOSE 8080

# Run
ENTRYPOINT task run
