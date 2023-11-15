package qa

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generationCommand)
}

var generationCommand = &cobra.Command{
	Use:     "generate",
	Short:   "Generation an application or a code snippet",
	Aliases: []string{"g"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generation of the application")
		fmt.Println(args)

		// what do we do with args
	},
}
