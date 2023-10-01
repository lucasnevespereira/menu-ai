package main

import (
	"context"
	"log"
	"menu-ai/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type MenuResponse struct {
	Completion string `json:"completion"`
}

type MenuRequest struct {
	Prompt string `json:"prompt"`
}

func getOpenAiCompletion(prompt string) (string, error) {
	conf := configs.Load()
	client := openai.NewClient(conf.OpenAiApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		log.Printf("getOpenAiCompletion: %v \n", err.Error())
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func generateMenuHandler(c *gin.Context) {
	var req MenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	completion, err := getOpenAiCompletion(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := MenuResponse{Completion: completion}
	c.JSON(http.StatusOK, resp)
}

func main() {
	r := gin.Default()
	r.Static("/", "./ui")

	r.POST("/generate-menu", generateMenuHandler)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
