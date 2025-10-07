# grpc
Repository for trying out grpc

## Install Protocol Buffer Compiler Installation
```sh
brew install protobuf
```

## Install `protoc-gen-go` and `protoc-gen-go-grpc` plugins
```sh
brew install protoc-gen-go protoc-gen-go-grpc
```

## Generate Go code from the .proto
```sh
cd demo-grpc
make build
```

## Update Go module
```sh
go mod tidy
```

## Set GOROOT environment variable
```sh
export GOROOT=/opt/homebrew/opt/go/libexec
```


## Trouble shooting
### Check what's using gRPC port (50051)
```sh
lsof -i :50051
```

### Kill the specified process by PID
```sh
kill {PID}
```
