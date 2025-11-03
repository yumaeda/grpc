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

### Install `protoc-gen-grpc-federation`
```sh
go install github.com/mercari/grpc-federation/cmd/protoc-gen-grpc-federation@latest
```

## Build
### Generate Go code from the .proto
```sh
make build
```

## Run Server
```sh
export GOROOT=/opt/homebrew/opt/go/libexec
export TIDB_CONFIG_JSON='{"db.password":"{your_password}","db.host":"{your_host}","db.name":"{your_db}","db.user":"{your_user}"}'
make run_server
```

## Run Tests
### In another terminal, run the tests
```sh
./test_grpc.sh
```

## gRPC Federation
### Build
#### Generate Go code from the .proto
```sh
make build_federation
```

### Run
#### Run all the gRPC servers
```sh
export GOROOT=/opt/homebrew/opt/go/libexec
export TIDB_CONFIG_JSON='{"db.password":"{your_password}","db.host":"{your_host}","db.name":"{your_db}","db.user":"{your_user}"}'
make run_all_servers
```

### Run tests
### In another terminal, run the tests
```sh
./test_grpc_federation.sh
```

## Trouble shooting
### Update Go module
```sh
go mod tidy
```

### Check what's using gRPC ports
```sh
lsof -i :50050  # BFF server
lsof -i :50051  # restaurant server
lsof -i :50052  # video server
lsof -i :50053  # swapi server
```

### Kill the specified process by PID
```sh
kill {PID}
```

## Reference
- https://medium.com/@rashid14713524/introduction-to-grpc-using-golang-for-beginners-e37d7dbafef2
- https://qiita.com/haruotsu/items/762e7b292dddaf6250cb
