package main

import (
	"log"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	restaurantpb "github.com/yumaeda/grpc/swapi/restaurant/restaurant"
	swapipb "github.com/yumaeda/grpc/swapi/swapi"
	videopb "github.com/yumaeda/grpc/swapi/video/video"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type client struct{}

func (c *client) Swapi_Restaurant_RestaurantServiceClient(_ swapipb.SWAPIClientConfig) (restaurantpb.RestaurantServiceClient, error) {
	ep := os.Getenv("RESTAURANT_SERVICE_ENDPOINT")
	conn, err := grpc.NewClient(ep,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	)
	if err != nil {
		return nil, err
	}
	return restaurantpb.NewRestaurantServiceClient(conn), nil
}

func (c *client) Swapi_Video_VideoServiceClient(_ swapipb.SWAPIClientConfig) (videopb.VideoServiceClient, error) {
	ep := os.Getenv("VIDEO_SERVICE_ENDPOINT")
	conn, err := grpc.NewClient(ep,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	)
	if err != nil {
		return nil, err
	}
	return videopb.NewVideoServiceClient(conn), nil
}

func run() error {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		return err
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	server, err := swapipb.NewSWAPI(swapipb.SWAPIConfig{
		Client: new(client),
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	})
	if err != nil {
		return err
	}

	swapipb.RegisterSWAPIServer(grpcServer, server)
	reflection.Register(grpcServer)

	log.Println("gRPC server listening on :50053")
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
