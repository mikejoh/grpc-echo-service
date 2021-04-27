# Build
FROM golang:1.15 AS build
ADD . /grpc-echo-service/
WORKDIR /grpc-echo-service/
RUN go build -o echo_server/echo_server ./echo_server/*.go
ENTRYPOINT ["./echo_server/echo_server"]