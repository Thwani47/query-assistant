package qa

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
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
		client := openai.NewClient("") // TODO find a way to pass the key here (configure the CLI). also move this to a seperate file
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
