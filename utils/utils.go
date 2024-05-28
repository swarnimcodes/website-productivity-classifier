package utils

type GeminiModel struct {
	Name                       string   `json:"name"`
	Version                    string   `json:"version"`
	DisplayName                string   `json:"displayName"`
	Description                string   `json:"description"`
	InputTokenLimit            int      `json:"inputTokenLimit"`
	OutputTokenLimit           int      `json:"outputTokenLimit"`
	SupportedGenerationMethods []string `json:"supportedGenerationMethods"`
	Temperature                float64  `json:"temperature"`
	TopP                       float64  `json:"topP"`
	TopK                       float64  `json:"topK"`
}

type OpenAiModel struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	OwnedBy string `json:"owned_by"`
}

type Models struct {
	Object string        `json:"object"`
	Data   []OpenAiModel `json:"data"`
}

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletion struct {
	Model       string                  `json:"model"`
	Messages    []ChatCompletionMessage `json:"messages"`
	Temperature float64                 `json:"temperature"`
}

type ResponseChoice struct {
	Index        int                   `json:"index"`
	Message      ChatCompletionMessage `json:"message"`
	LogProbs     string                `json:"logProbs"`
	FinishReason string                `json:"finishReason"`
}

type ResponseUsage struct {
	PromptTokens     int `json:"PromptTokens"`
	CompletionTokens int `json:"CompletionTokens"`
	TotalTokens      int `json:"TotalTokens"`
}

type ChatResponse struct {
	Id                string           `json:"id"`
	Object            string           `json:"object"`
	Created           int              `json:"created"`
	Model             string           `json:"model"`
	Choices           []ResponseChoice `json:"choices"`
	Usage             ResponseUsage    `json:"usage"`
	SystemFingerprint string           `json:"systemFingerprint"`
}

type InputMessage struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}
