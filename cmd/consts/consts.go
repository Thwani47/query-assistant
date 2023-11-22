package consts

type PromptType int

const (
	TextPrompt     PromptType = 0
	PasswordPrompt PromptType = 1
	SelectPrompt   PromptType = 2
	OpenAI_APIKey  string     = "OPENAI_API_KEY"
)

type PromptItem struct {
	ID            string
	Label         string
	Value         string
	SelectOptions []string
	PromptType    PromptType
}
