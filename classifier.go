package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/swarnimcodes/website-productivity-classifier/openai"
	"github.com/swarnimcodes/website-productivity-classifier/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	openAiBaseUrl := os.Getenv("BASE_URL")
	openAiApiKey := os.Getenv("OPENAI_API_KEY")
	fmt.Println(openAiApiKey)
	models, err := openai.RetrieveAvailableModels(openAiBaseUrl, openAiApiKey)
	if err != nil {
		log.Fatalf("Could not retrieve OpenAI Models: %v", err)
	}

	for _, model := range models.Data {
		fmt.Printf("ID: %s\n", model.Id)
		fmt.Printf("Object: %s\n", model.Object)
		fmt.Printf("Created: %d\n", model.Created)
		fmt.Printf("Owned By: %s\n", model.OwnedBy)
		fmt.Println()
	}

	context := `
	You are a website productivity grading system.
	You will receive a website url and its html description tag like so:

	{
		"url": "https://platform.openai.com/docs/overview",
		"description": "Explore resources, tutorials, API docs, and 
		dynamic examples to get the most out of OpenAI's developer platform."
	}

	Your job is to classify the website as productive, unproductive or neutral
	1 being productive
	2 being unproductive
	3 being neutral
	
	So your output needs to be of the following format:
	{
		"url": "https://platform.openai.com/docs/overview",
		"productivityName": "productive"
		"productivity": 1
	}
	`

	c := utils.ChatCompletion{
		Model:       "gpt-3.5-turbo-0613",
		Temperature: 0.2,
		Messages: []utils.ChatCompletionMessage{
			{
				Role:    "system",
				Content: context,
			},
		},
	}

	input := utils.InputMessage{
		URL: "https://www.practical-go-lessons.com/",
		Description: `
		My main objective is to teach you the language 
		in a progressive way. I also tried to clarify and explain 
		some common computer science notions that can be difficult 
		to grasp, especially for newcomers.
		`,
	}
	jsonInput, err := json.Marshal(input)

	// TODO: handle error

	c.Messages = append(c.Messages, utils.ChatCompletionMessage{
		Role:    "user",
		Content: string(jsonInput),
	})

	newc, err := openai.Chat(c)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	responseMessage := newc.Choices[0].Message.Content
	jsonNewc, _ := json.MarshalIndent(newc, "", "  ")
	fmt.Println(string(jsonNewc))
	fmt.Println(responseMessage)
}
