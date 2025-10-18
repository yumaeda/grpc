# grpc
Repository for trying out grpc

## Preparation
### Install Protocol Buffer Compiler Installation
```sh
brew install protobuf
```

### Install `protoc-gen-go` and `protoc-gen-go-grpc` plugins
```sh
brew install protoc-gen-go protoc-gen-go-grpc
```

### Install `grpcurl`
```sh
brew install grpcurl
```

## Build
### Move to the below directory
```sh
cd demo-grpc
```

### Generate Go code from the .proto
```sh
make build
```

## Run Server
### Set GOROOT environment variable
```sh
export GOROOT=/opt/homebrew/opt/go/libexec
make run_server
```

## Run Client
### Set GOROOT environment variable
```sh
export GOROOT=/opt/homebrew/opt/go/libexec
make run_client
```

## Trouble shooting
### Update Go module
```sh
go mod tidy
```

### Check what's using gRPC port (50051)
```sh
lsof -i :50051
```

### Kill the specified process by PID
```sh
kill {PID}
```

## Reference
- https://medium.com/@rashid14713524/introduction-to-grpc-using-golang-for-beginners-e37d7dbafef2
