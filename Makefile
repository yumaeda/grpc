build:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	proto/admin_user/admin_user.proto \
	proto/area/area.proto \
	proto/category/category.proto \
	proto/dish/dish.proto \
	proto/drink/drink.proto \
	proto/menu/menu.proto \
	proto/photo/photo.proto \
	proto/ranking/ranking.proto \
	proto/genre/genre.proto

PROTOC := PATH=$$PATH:$$(go env GOPATH)/bin protoc
PROTO_PATH := ./proto
GRPC_FED_PROTO_PATH := $$(go list -m -f '{{.Dir}}' github.com/mercari/grpc-federation)/proto
GOOGLEAPIS_PROTO_PATH := $$(go list -m -f '{{.Dir}}' github.com/mercari/grpc-federation)/proto_deps
GO_OPTS := --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

build_federation:
	@mkdir -p swapi/swapi swapi/restaurant swapi/video
	$(PROTOC) --proto_path=$(PROTO_PATH) \
		--go_out=swapi/restaurant $(GO_OPTS) \
		--go-grpc_out=swapi/restaurant \
		proto/restaurant/restaurant.proto
	$(PROTOC) --proto_path=$(PROTO_PATH) \
		--go_out=swapi/video $(GO_OPTS) \
		--go-grpc_out=swapi/video \
		proto/video/video.proto
	$(PROTOC) --proto_path=$(PROTO_PATH) \
		--proto_path=$(GRPC_FED_PROTO_PATH) \
		--proto_path=$(GOOGLEAPIS_PROTO_PATH) \
		--go_out=swapi $(GO_OPTS) \
		--go-grpc_out=swapi \
		--grpc-federation_out=swapi \
		--grpc-federation_opt=paths=source_relative \
		./proto/swapi.proto
	@mv swapi/*.pb.go swapi/swapi/

run_server:
	@GOROOT="" go run cmd/bff/main.go

run_restaurant_server:
	@echo "Starting restaurant server on port 50051..."
	@GOROOT="" go run cmd/restaurant/main.go &

run_video_server:
	@echo "Starting video server on port 50052..."
	@GOROOT="" go run cmd/video/main.go &

run_swapi_server:
	@echo "Starting federation server on port 50053..."
	@RESTAURANT_SERVICE_ENDPOINT=localhost:50051 \
	VIDEO_SERVICE_ENDPOINT=localhost:50052 \
	GOROOT="" go run cmd/swapi/main.go &

run_all_servers:
	@echo "Starting all servers..."
	@GOROOT="" go run cmd/restaurant/main.go &
	@GOROOT="" go run cmd/video/main.go &
	@sleep 2  # Wait for backend servers to start
	@RESTAURANT_SERVICE_ENDPOINT=localhost:50051 \
	VIDEO_SERVICE_ENDPOINT=localhost:50052 \
	GOROOT="" go run cmd/swapi/main.go &
	@echo "All servers started in background"

stop_servers:
	@echo "Stopping servers on ports 50051, 50052, 50053..."
	@for port in 50051 50052 50053; do \
		if lsof -ti :$$port > /dev/null 2>&1; then \
			echo "  Port $$port: stopping (PID: $$(lsof -ti :$$port))"; \
			lsof -ti :$$port | xargs kill -9 2>/dev/null; \
		else \
			echo "  Port $$port: not in use"; \
		fi \
	done
	@echo "Done"

clean:
	rm -f proto/*/*.pb.go
	rm -rf swapi/
