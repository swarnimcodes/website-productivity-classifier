package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/swarnimcodes/website-productivity-classifier/utils"
)

func Chat(chatCompletion utils.ChatCompletion) (utils.ChatResponse, error) {
	endpoint := "v1/chat/completions"
	openaiBaseUrl := os.Getenv("BASE_URL")
	openAiApiKey := os.Getenv("OPENAI_API_KEY")
	url := fmt.Sprintf("%s/%s", openaiBaseUrl, endpoint)
	bearerToken := fmt.Sprintf("Bearer %s", openAiApiKey)

	marshalled, err := json.Marshal(chatCompletion)
	if err != nil {
		return utils.ChatResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(marshalled))
	if err != nil {
		return utils.ChatResponse{}, err
	}
	req.Header.Set("Authorization", bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return utils.ChatResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		err := fmt.Sprintf("Non-2xx status code returned: %d", resp.StatusCode)
		return utils.ChatResponse{}, errors.New(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.ChatResponse{}, err
	}

	var c utils.ChatResponse
	err = json.Unmarshal([]byte(body), &c)
	if err != nil {
		return utils.ChatResponse{}, err
	}
	return c, nil
}
