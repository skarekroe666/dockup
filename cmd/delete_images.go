package cmd

import (
	"fmt"
	"log"
	"strings"

	imagetypes "github.com/docker/docker/api/types/image"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var removeAll bool

var removeImgCmd = &cobra.Command{
	Use:   "rmi",
	Short: "remove the images",
	Run: func(cmd *cobra.Command, args []string) {
		removeImage()
	},
}

func removeImage() {
	ctx, cli := dockerClient()
	defer cli.Close()

	images, err := cli.ImageList(ctx, imagetypes.ListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}

	if len(images) == 0 {
		fmt.Println("No images present")
		return
	}

	//if the flag is set to -a delete all the images
	if removeAll {
		prompt := promptui.Select{
			Label: "Are you sure you want to remove all the images",
			Items: []string{"Yes", "No"},
		}

		_, option, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}

		if option == "Yes" {
			for _, img := range images {
				fmt.Printf("Deleting image %s (%s)...\n", img.RepoDigests[0][1:], img.ID[:7])
				_, err = cli.ImageRemove(ctx, img.ID, imagetypes.RemoveOptions{})
				if err != nil {
					log.Fatal(err)
				}
			}
			fmt.Println("Deleted all images.")
			return
		} else {
			fmt.Println("Images not removed")
			return
		}
	}

	for _, img := range images {
		listOfImages = append(listOfImages, img.RepoDigests[0][:1]+" - "+img.ID)
	}

	prompt := promptui.Select{
		Label: "Select an image to delete",
		Items: listOfImages,
	}

	_, selection, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	prompt = promptui.Select{
		Label: "Are you sure you want ot delete " + selection + " ?",
		Items: []string{"Yes", "No"},
	}

	_, option, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if option == "Yes" {
		split := strings.Split(selection, " - ")

		_, err = cli.ImageRemove(ctx, split[1], imagetypes.RemoveOptions{Force: true})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Image %s deleted successfully\n", split[0])
	} else {
		fmt.Println("No Image deleted")
		return
	}
}

func init() {
	imagesCmd.AddCommand(removeImgCmd)
	removeImgCmd.Flags().BoolVarP(&removeAll, "all", "a", false, "Delete all images")
}
