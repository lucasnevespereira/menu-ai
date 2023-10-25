package main

import (
	"fmt"
	"log"
	"menu-ai/configs"
	"menu-ai/internal/connectors/openai"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuResponse struct {
	Completion string `json:"completion"`
}

type MenuRequest struct {
	Prompt string `json:"prompt"`
}

func getOpenAiCompletion(prompt string) (string, error) {
	conf := configs.Load()
	log.Printf("prompt: %s \n", prompt)
	openAIClient := openai.NewClient(conf.OpenAiApiKey)
	content, err := openAIClient.GetCompletion(prompt)
	if err != nil {
		log.Printf("getOpenAiCompletion: %v \n", err.Error())
		return "", err
	}
	fmt.Printf("content: %v", content)
	return content, nil
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
