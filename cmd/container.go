package cmd

import (
	"fmt"
	"log"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/spf13/cobra"
)

var listOfContainers = []string{}

var ContainerCmd = &cobra.Command{
	Use:   "container [command] [flags]",
	Short: "Use this command to list all the containers",
	Run: func(cmd *cobra.Command, args []string) {
		listContainer()
	},
}

func listContainer() {
	ctx, cli := dockerClient()
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}

	for _, container := range containers {
		fmt.Printf("%s\t%s\n", container.ID[:12], container.Image)
	}
}

func init() {
	rootCmd.AddCommand(ContainerCmd)
}
