package cmd

import (
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "use this command to rename a container",
	Run: func(cmd *cobra.Command, args []string) {
		// renameContaier()
	},
}

// func renameContaier() {
// 	ctx, cli := dockerClient()
// 	containerList := listContainer(ctx, cli)

// 	if len(containerList) == 0 {
// 		fmt.Println("No Containers running")
// 		return
// 	}

// 	for _, container := range containerList {
// 		runningContainer = append(runningContainer, container.Names[0][1:]+" - "+container.ID[:6])
// 	}

// 	fmt.Println("CONTAINER NAME - CONTAINER ID")
// 	prompt := promptui.Select{
// 		Label: "Select a container to rename",
// 		Items: runningContainer,
// 	}

// 	_, result, err := prompt.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Print("Enter new name: ")
// 	scanner.Scan()
// 	newName := scanner.Text()
// 	newName = strings.TrimSpace(newName)

// 	if newName == "" {
// 		fmt.Println("Name cannot be empty")
// 		return
// 	}

// 	split := strings.Split(result, " - ")

// 	err = cli.ContainerRename(ctx, split[0], newName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func init() {
	// ContainerCmd.AddCommand(renameCmd)
	rootCmd.AddCommand(renameCmd)
}
