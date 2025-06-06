package inference

import (
	"context"
	"fmt"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/honganh1206/clue/conversation"
	"github.com/honganh1206/clue/tools"
)

type Model interface {
	// FIXME: VERY RESOURCE-CONSUMING since we are invoking this in every loop
	// What to do? Maintain a parallel flattened view/Flatten incrementally with new messages/Modify the engine
	RunInference(ctx context.Context, msgs []*conversation.MessageParam, tools []tools.ToolDefinition) (*conversation.MessageResponse, error)
	Name() string
}

type ModelConfig struct {
	Provider   string
	PromptPath string
	Model      string
	MaxTokens  int64
	ListModel  bool
}

func Init(config ModelConfig) (Model, error) {
	switch config.Provider {
	case AnthropicProvider:
		client := anthropic.NewClient() // Default to look up ANTHROPIC_API_KEY
		return NewAnthropicModel(&client, config.PromptPath, ModelVersion(config.Model), config.MaxTokens), nil
	default:
		return nil, fmt.Errorf("unknown model provider: %s", config.Provider)
	}
}

func ListAvailableModels(provider ProviderName) []ModelVersion {
	switch provider {
	case AnthropicProvider:
		return []ModelVersion{
			Claude4Opus,
			Claude4Sonnet,
			Claude37Sonnet,
			Claude35Sonnet,
			Claude35Haiku,
			Claude3Opus,
			Claude3Sonnet, // FIXME: Deprecated soon
			Claude3Haiku,
		}
	default:
		return []ModelVersion{}
	}
}

func GetDefaultModel(provider ProviderName) ModelVersion {
	switch provider {
	case AnthropicProvider:
		return ModelVersion(anthropic.ModelClaude4Sonnet20250514)
	default:
		return ""
	}
}

// formatModelsForHelp formats a list of models for help text
func FormatModelsForHelp(models []ModelVersion) string {
	if len(models) == 0 {
		return ""
	}

	modelStrings := make([]string, len(models))
	for i, model := range models {
		modelStrings[i] = string(model)
	}
	return strings.Join(modelStrings, ", ")
}
