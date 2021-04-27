# Build
FROM golang:1.15 AS build
ADD . /echo_server/
COPY go.mod /echo_server/
WORKDIR /echo_server/
RUN go build -o echo_server .

# Final
FROM gcr.io/distroless/base-debian10
COPY --from=build /echo_server/echo_server /bin/
ENTRYPOINT ["echo_server"]