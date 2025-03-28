FROM golang:1.24-alpine

# Build optimization
ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

# Copy source code
COPY . .

# Build the application
RUN go build -o /bin/notification ./services/notification/cmd

# Set port and startup command
EXPOSE 8082
CMD ["/bin/notification"] 