FROM golang:1.24-alpine

# Tối ưu hóa cho build
ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

# Copy toàn bộ source code
COPY . .

# Build the application
RUN go build -o /bin/gateway ./services/gateway/cmd

# Thiết lập port và lệnh khởi động
EXPOSE 8080
CMD ["/bin/gateway"] 