package cmd

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

func dockerClient() (context.Context, *client.Client) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	return ctx, cli
}

func listContainer(ctx context.Context, cli *client.Client) []container.Summary {
	containerList, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	return containerList
}
