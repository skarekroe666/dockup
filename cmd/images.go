package cmd

import (
	"fmt"
	"log"

	"github.com/docker/docker/api/types/image"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "image",
	Short: "Use this command to list all the images",
	Run: func(cmd *cobra.Command, args []string) {
		listImage()
	},
}

func listImage() {
	ctx, cli := dockerClient()
	images, err := cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}
