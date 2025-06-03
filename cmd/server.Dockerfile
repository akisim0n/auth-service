FROM golang:1.24-alpine3.21 AS builder

COPY . /auth-server/
WORKDIR /auth-server/

RUN go mod download
RUN go build -o ./bin/auth-server server/main.go

FROM alpine:3.21

WORKDIR /root/

COPY --from=builder /auth-server/bin/auth-server .

CMD ["./auth-server"]