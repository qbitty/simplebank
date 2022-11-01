FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
ENV GOPROXY https://goproxy.cn,direct
RUN go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]
