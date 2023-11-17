package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of Query Assistant",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Query Assistant v0.0.1")
	},
}
