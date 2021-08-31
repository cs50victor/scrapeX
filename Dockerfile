FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \ 
    PORT=3000

# Move to working directory /app
WORKDIR /app

# Download necessary Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy all current files into the container
COPY . .

# Build the application with dafault name - scrapper
RUN go build

# Build a small image
FROM scratch

COPY --from=builder /app/scrapper /
COPY --from=builder /app/stockTrend /
COPY --from=builder /app/5am.json /

# Export necessary port
EXPOSE $PORT

# Command to run when starting the container
ENTRYPOINT ["./scrapper"]