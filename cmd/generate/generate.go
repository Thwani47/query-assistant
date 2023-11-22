package generate

import (
	"context"
	"fmt"

	"github.com/Thwani47/query-assistant/cmd/consts"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {

}

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate an application or a code snippet",
	Aliases: []string{"g"},
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString(consts.OpenAI_APIKey)
		fmt.Println(apiKey)
		client := openai.NewClient(viper.GetString(consts.OpenAI_APIKey))
		response, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: args[0],
				},
			},
		},
		)

		if err != nil {
			fmt.Printf("Chat completion error: %v\n", err)
		}

		fmt.Println(response.Choices[0].Message.Content)

	},
}
