FROM golang:1.21
WORKDIR /app
COPY fib.go .
RUN go build -o fib fib.go
CMD ["./fib"]