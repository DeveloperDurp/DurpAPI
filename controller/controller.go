package controller

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

type Controller struct {
	openaiClient *openai.Client
	unraidAPIKey string
	unraidURI    string
}

func NewController() *Controller {
	err := godotenv.Load(".env")

	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	openaiClient := openai.NewClient(openaiApiKey)
	unraidAPIKey := os.Getenv("UNRAID_API_KEY")
	unraidURI := os.Getenv("UNRAID_URI")

	if err != nil {
		fmt.Println(".env file not found, using environment variables")
	}
	return &Controller{
		openaiClient: openaiClient,
		unraidAPIKey: unraidAPIKey,
		unraidURI:    unraidURI,
	}
}

type Message struct {
	Message string `json:"message" example:"message"`
}
