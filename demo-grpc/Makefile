build:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative check_prime/check_prime.proto

run_server:
	go run server/server.go

run_client:
	go run client/client.go

clean:
	rm -f check_prime/*.pb.go
