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

## Run Client
```sh
export GOROOT=/opt/homebrew/opt/go/libexec
grpcurl -plaintext -d '{"id":1}' localhost:50051 area.AreaService/GetArea 
grpcurl -plaintext -d '{"id":"0b95fe0d-a323-43d7-954b-2164a1d3242d"}' localhost:50051 restaurant.RestaurantService/GetRestaurant
grpcurl -plaintext -d '{"id":"116e70bb-c26c-4ec7-8935-7f922e8bf551"}' localhost:50051 menu.MenuService/GetMenu
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
- https://qiita.com/haruotsu/items/762e7b292dddaf6250cb
