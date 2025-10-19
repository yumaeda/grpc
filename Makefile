build:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative internal/proto/area/area.proto internal/proto/restaurant/restaurant.proto

run_server:
	go run cmd/server/main.go

clean:
	rm -f internal/proto/area/*.pb.go internal/proto/restaurant/*.pb.go
