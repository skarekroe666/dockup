package cmd

import (
	"fmt"
	"log"
	"strings"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var deleteAll bool

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a container",
	Run: func(cmd *cobra.Command, args []string) {
		deleteContainer()
	},
}

func deleteContainer() {
	ctx, cli := dockerClient()
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}

	if len(containers) == 0 {
		fmt.Print("No container running/present")
		return
	}

	//if the flag is set to -a delete all the containers
	if deleteAll {

		prompt := promptui.Select{
			Label: "Are you sure you want to delete all the containers?",
			Items: []string{"Yes", "No"},
		}

		_, option, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}

		if option == "Yes" {
			for _, container := range containers {
				fmt.Printf("Deleting container %s (%s)...\n", container.Names[0][1:], container.ID[:6])
				err := cli.ContainerRemove(ctx, container.ID, containertypes.RemoveOptions{Force: true})
				if err != nil {
					log.Fatal(err)
				}
			}
			fmt.Println("Deleted all containers.")
			return
		} else {
			fmt.Println("Containers not removed")
			return
		}
	}

	//if the flag is not set to -a/ to delete a specific container
	for _, container := range containers {
		listOfContainers = append(listOfContainers, container.Names[0][1:]+" - "+container.ID)
	}

	prompt := promptui.Select{
		Label: "Select a container to delete",
		Items: listOfContainers,
	}

	_, selection, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	prompt = promptui.Select{
		Label: "Are you sure you want ot delete " + selection[:7] + " ?",
		Items: []string{"Yes", "No"},
	}

	_, option, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if option == "Yes" {
		split := strings.Split(selection, " - ")

		err := cli.ContainerRemove(ctx, split[1], containertypes.RemoveOptions{Force: true})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Container %s deleted successfully\n", split[0])
	} else {
		fmt.Println("No Container deleted")
		return
	}
}

func init() {
	ContainerCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&deleteAll, "a", "", false, "Delete all containers")
}
