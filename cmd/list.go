package cmd

import (
	"context"
	"log"

	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

func dockerClient() (context.Context, *client.Client) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithHost("unix:///home/skarekroe/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	return ctx, cli
}
