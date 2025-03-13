package provider

import (
	"errors"
	"os"

	"github.com/Gustrb/kramer/internal/repository"
)

// AssistantProvider is any type that can provide AI assistant services, such as
// Chatgpt, etc..
type AssistantProvider interface {
	// GetResponse returns a response from the AI assistant. It should be context-aware
	// and return a response based on the inquiry.
	GetResponse(id int, inquiry string) (string, error)

	// LoadContext loads the context with the given ID
	LoadContext(id int) error
}

type Provider string

const (
	ChatGPT Provider = "chatgpt"
)

var (
	ErrChatGPTAPIKeyNotSet  = errors.New("chatgpt api key not set")
	ErrProviderNotSupported = errors.New("provider not supported")
	ErrChatGPTInvalidModel  = errors.New("chatgpt invalid model")
)

func ProviderFactory(store repository.Store, provider Provider) (AssistantProvider, error) {
	switch provider {
	case ChatGPT:
		// Get the key from the environment
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			return nil, ErrChatGPTAPIKeyNotSet
		}

		model := ChatGPT4oMiniModel
		m := os.Getenv("OPENAI_MODEL")
		if m != "" {
			if IsValidModel(Model(m)) {
				model = Model(m)
			} else {
				return nil, ErrChatGPTInvalidModel
			}
		}

		return NewChatGPTAssistantProvider(store, model, apiKey), nil

	default:
		return nil, ErrProviderNotSupported
	}
}
