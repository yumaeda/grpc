package main

import (
	"log"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	photopb "github.com/yumaeda/sakabas-grpc/swapi/photo/photo"
	restaurantpb "github.com/yumaeda/sakabas-grpc/swapi/restaurant/restaurant"
	swapipb "github.com/yumaeda/sakabas-grpc/swapi/swapi"
	videopb "github.com/yumaeda/sakabas-grpc/swapi/video/video"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type client struct {
	restaurantConn *grpc.ClientConn
	videoConn      *grpc.ClientConn
	photoConn      *grpc.ClientConn
}

func (c *client) Close() error {
	if c.restaurantConn != nil {
		if err := c.restaurantConn.Close(); err != nil {
			return err
		}
	}
	if c.videoConn != nil {
		if err := c.videoConn.Close(); err != nil {
			return err
		}
	}
	if c.photoConn != nil {
		if err := c.photoConn.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) Swapi_Restaurant_RestaurantServiceClient(_ swapipb.SWAPIClientConfig) (restaurantpb.RestaurantServiceClient, error) {
	if c.restaurantConn == nil {
		ep := os.Getenv("RESTAURANT_SERVICE_ENDPOINT")
		conn, err := grpc.NewClient(ep,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		)
		if err != nil {
			return nil, err
		}
		c.restaurantConn = conn
	}
	return restaurantpb.NewRestaurantServiceClient(c.restaurantConn), nil
}

func (c *client) Swapi_Video_VideoServiceClient(_ swapipb.SWAPIClientConfig) (videopb.VideoServiceClient, error) {
	if c.videoConn == nil {
		ep := os.Getenv("VIDEO_SERVICE_ENDPOINT")
		conn, err := grpc.NewClient(ep,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		)
		if err != nil {
			return nil, err
		}
		c.videoConn = conn
	}
	return videopb.NewVideoServiceClient(c.videoConn), nil
}

func (c *client) Swapi_Photo_PhotoServiceClient(_ swapipb.SWAPIClientConfig) (photopb.PhotoServiceClient, error) {
	if c.photoConn == nil {
		ep := os.Getenv("PHOTO_SERVICE_ENDPOINT")
		conn, err := grpc.NewClient(ep,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		)
		if err != nil {
			return nil, err
		}
		c.photoConn = conn
	}
	return photopb.NewPhotoServiceClient(c.photoConn), nil
}

func run() error {
	listener, err := net.Listen("tcp", ":50054")
	if err != nil {
		return err
	}
	defer listener.Close()

	// Create client and ensure connections are closed on exit
	cli := new(client)
	defer cli.Close()

	grpcServer := grpc.NewServer()
	server, err := swapipb.NewSWAPI(swapipb.SWAPIConfig{
		Client: cli,
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	})
	if err != nil {
		return err
	}

	swapipb.RegisterSWAPIServer(grpcServer, server)
	reflection.Register(grpcServer)

	log.Println("gRPC server listening on :50054")
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
