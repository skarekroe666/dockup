package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

func loadEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("couldn't load .env file: %w", err)
	}

	hostName := os.Getenv("DOCKER_HOST")

	return hostName, nil
}

func dockerClient() (context.Context, *client.Client) {
	host, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithHost(host),
		client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	return ctx, cli
}
