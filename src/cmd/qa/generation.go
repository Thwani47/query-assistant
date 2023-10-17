package qa

import "github.com/spf13/cobra"

func init() {

}

var generationCommand = &cobra.Command{
	Use:   "generation",
	Short: "Generation of the application",
}
