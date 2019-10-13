FROM golang:1.12 AS builder
WORKDIR /module
COPY . /module
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:3  
WORKDIR /root/
COPY --from=builder /module/main /root
COPY ui /root/ui
CMD ["./main"] 