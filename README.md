# The simplest gRPC example

_Heavily inspired by the official grpc [`helloworld`](https://grpc.io/docs/languages/go/quickstart/) example._

## How to
1. Download `protoc`
2. In the root of this repository run:
```
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative echo/echo.prot
```
3. Run the server
4. Run the client