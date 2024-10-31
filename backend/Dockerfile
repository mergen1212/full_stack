FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:latest  
WORKDIR /root
COPY --from=builder /app/main .
CMD ["/root/main"]