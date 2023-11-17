package cmd

import (
	"fmt"
	"os"

	"github.com/Thwani47/query-assistant/cmd/generate"
	"github.com/Thwani47/query-assistant/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qa",
	Short: "Query Assistant is your AI-based terminal companion",
	Long: `
Query Assistant is a command-line tool designed to be your digital assistant. 
With the power of AI, it can help you streamline everyday tasks and answer your questions. 
Whether you need to create code, launch applications, stop music playback, retrieve upcoming meeting details, or tackle various other tasks, Query Assistant is here to assist. 
Simply type 'qa' in your terminal to engage with your digital helper. 
Empower your command line with the capabilities of Query Assistant and make your daily computing experience more efficient and user-friendly.`,
}

func init() {
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(version.VersionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
