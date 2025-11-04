package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var run = &cobra.Command{
	Use:   "run",
	Short: "run the container",
	Run: func(cmd *cobra.Command, args []string) {
		RunContainer()
	},
}

func init() {
	rootCmd.AddCommand(run)
}

func RunContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}
}
