package provider

import (
	"context"

	"github.com/Gustrb/kramer/internal/repository"
	"github.com/Gustrb/kramer/models"
)

const apiURL = "https://api.openai.com/v1/chat/completions"

type Role string
type Model string

const (
	// ChatGPTSystemRole is the role of the system in the conversation.
	ChatGPTSystemRole Role = "system"

	// ChatGPTUserRole is the role of the user in the conversation.
	ChatGPTUserRole Role = "user"
)

const (
	ChatGPT4BModel     Model = "gpt-4"
	ChatGPT4oMiniModel Model = "gpt-4o-mini"
)

func IsValidModel(model Model) bool {
	switch model {
	case ChatGPT4BModel, ChatGPT4oMiniModel:
		return true
	default:
		return false
	}
}

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type ChatGPTAssistantProvider struct {
	store   repository.Store
	context map[int][]Message
	model   Model
	key     string
}

func NewChatGPTAssistantProvider(store repository.Store, model Model, apiKey string) *ChatGPTAssistantProvider {
	return &ChatGPTAssistantProvider{
		context: make(map[int][]Message),
		model:   model,
		key:     apiKey,
		store:   store,
	}
}

func (c *ChatGPTAssistantProvider) GetResponse(id int, inquiry string) (string, error) {
	// First check if there is a context for the given id (topic)
	if _, ok := c.context[id]; !ok {
		c.context[id] = []Message{}
	}

	// Append the user's inquiry to the context
	c.context[id] = append(c.context[id], Message{Role: ChatGPTUserRole, Content: inquiry})

	request := OpenAICompletionsRequest{
		Model:    c.model,
		Messages: c.context[id],
	}

	chatResponse, err := CallOpenAICompatibleCompletionsAPI(apiURL, c.key, request)
	if err != nil {
		return "", err
	}

	var responseText string
	if len(chatResponse.Choices) > 0 {
		responseText = chatResponse.Choices[0].Message.Content
		c.context[id] = append(c.context[id], chatResponse.Choices[0].Message)
	}

	c.store.History().Create(
		context.Background(),
		models.CreateHistoryEntry{
			Message:   inquiry,
			Role:      string(ChatGPTUserRole),
			ContextID: id,
		},
		models.CreateHistoryEntry{
			Message:   responseText,
			Role:      string(ChatGPTSystemRole),
			ContextID: id,
		},
	)

	return responseText, nil
}

func (c *ChatGPTAssistantProvider) LoadContext(id int) error {
	history, err := c.store.History().ReadByContextID(context.Background(), id)
	if err != nil {
		return err
	}

	c.context[id] = []Message{}
	for _, h := range history {
		c.context[id] = append(c.context[id], Message{Role: Role(h.Role), Content: h.Message})
	}

	return nil
}
