package config

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

type promptContent struct {
	errorMessage string
	label        string
}

var configItems = map[string]string{"Open AI API Key": "OpenAI_APIKey"}

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure QA",
	Long:  `Configure the settings for QA. For example, you can set the default model to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		configItemSelectPc := promptContent{
			errorMessage: "Please select item to configure",
			label:        "Configuration Item",
		}

		configItem := prompForSelect(configItemSelectPc)

		configItemInputPc := promptContent{
			errorMessage: fmt.Sprintf("Please enter a valid %s", configItem),
			label:        configItem,
		}

		inputConfigItem := promptForInput(configItemInputPc)

		fmt.Printf("You chose %s for %s", inputConfigItem, configItems[inputConfigItem])
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	err := viper.WriteConfigAs("qa.config.yaml")

	if err != nil {
		fmt.Println(err)
	}

	ConfigCmd.PersistentFlags().StringVarP(&configFile, "config-file", "f", "", "config file (default is $HOME/.qa.config.yaml)")
	ConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".qa")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func prompForSelect(pc promptContent) string {
	prompt := promptui.Select{
		Label: pc.label,
		Items: []string{"Open AI API Key", "Quit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return result
}

func promptForInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return fmt.Errorf(pc.errorMessage)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Failed to read input for %s: %s", pc.label, err)
		os.Exit(1)
	}

	return result
}
