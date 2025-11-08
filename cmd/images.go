package cmd

import (
	"fmt"
	"log"

	"github.com/docker/docker/api/types/image"
	"github.com/spf13/cobra"
)

var listOfImages = []string{}

var imagesCmd = &cobra.Command{
	Use:   "image",
	Short: "Use this command to list all the images",
	Run: func(cmd *cobra.Command, args []string) {
		listImage()
	},
}

func listImage() {
	ctx, cli := dockerClient()
	defer cli.Close()

	images, err := cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		fmt.Printf("%s\t%s\n", image.ID[:12], image.RepoTags)
	}
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}
