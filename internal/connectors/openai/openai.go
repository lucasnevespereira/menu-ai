package openai

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

type OpenAI struct {
	client *resty.Client
	apiKey string
}

type ChatCompletion struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewClient(apiKey string) *OpenAI {
	return &OpenAI{client: resty.New(), apiKey: apiKey}
}

func (c *OpenAI) GetCompletion(prompt string) (string, error) {
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo-16k-0613",
		"messages": []map[string]interface{}{
			{"role": "user", "content": prompt},
		},
	}

	response, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.apiKey).
		SetBody(requestBody).
		Post(apiEndpoint)

	if err != nil {
		return "", errors.Wrap(err, "GetCompletion calling endpoint")
	}

	var chatCompletion ChatCompletion
	if err := json.Unmarshal(response.Body(), &chatCompletion); err != nil {
		return "", errors.Wrap(err, "Failed to unmarshal response")
	}

	if len(chatCompletion.Choices) > 0 {
		return chatCompletion.Choices[0].Message.Content, nil
	}

	return "", errors.New("No content in choices")
}
