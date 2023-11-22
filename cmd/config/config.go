package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/Thwani47/query-assistant/cmd/consts"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var configFileName string = "qa.config.yaml"

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure QA",
	Long:  `Configure the settings for QA. For example, you can set the default model to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		configItems := []*consts.PromptItem{
			{
				ID:    consts.OpenAI_APIKey,
				Label: "Open AI API Key",
			},
		}

		_, err := configureSettings("Configure QA", 0, configItems)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}

		for _, config := range configItems {
			if config.Value != "" {
				viper.Set(config.ID, config.Value)
			}
		}
		viper.WriteConfigAs(configFileName)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	err := viper.WriteConfigAs(configFileName)

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

		viper.SetConfigName("qa.config")
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")

	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No config file found. Creating one...")
	}
}

func configureSettings(promptLabel string, startingIndex int, items []*consts.PromptItem) (bool, error) {
	doneID := "Done"

	if len(items) > 0 && items[0].ID != doneID {
		items = append([]*consts.PromptItem{{ID: doneID, Label: "Done"}}, items...)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Label | cyan }}",
		Inactive: "  {{ .Label | cyan }}",
		Selected: "\U0001F336 {{ .Label | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:        promptLabel,
		Items:        items,
		Size:         4,
		Templates:    templates,
		HideSelected: true,
		CursorPos:    startingIndex,
	}

	selectedIndex, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false, err
	}

	selectedItem := items[selectedIndex]

	if selectedItem.ID == doneID {
		return true, nil
	}

	var promptResponse string

	if selectedItem.PromptType == consts.TextPrompt || selectedItem.PromptType == consts.PasswordPrompt {
		promptResponse, err = promptForInput(*selectedItem)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return false, err
		}

		items[selectedIndex].Value = promptResponse
	}

	if selectedItem.PromptType == consts.SelectPrompt {
		promptResponse, err = prompForSelect(*selectedItem)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return false, err
		}

		items[selectedIndex].Value = promptResponse
	}

	return configureSettings(promptLabel, selectedIndex, items)
}

func prompForSelect(item consts.PromptItem) (string, error) {
	prompt := promptui.Select{
		Label:        item.Label,
		Items:        item.SelectOptions,
		HideSelected: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

func promptForInput(item consts.PromptItem) (string, error) {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("error: input cannot be empty")
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
		Label:       item.Label,
		Templates:   templates,
		Validate:    validate,
		HideEntered: true,
	}

	if item.PromptType == consts.PasswordPrompt {
		prompt.Mask = '*'
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed:%v\n", err)
		return "", err
	}

	return result, nil
}
