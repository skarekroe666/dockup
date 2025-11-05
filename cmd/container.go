package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var runningContainer = []string{}

var ContainerCmd = &cobra.Command{
	Use:   "container [command] [flags]",
	Short: "run the container",
	Run: func(cmd *cobra.Command, args []string) {
		// RunContainer()
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ContainerCmd)
}
