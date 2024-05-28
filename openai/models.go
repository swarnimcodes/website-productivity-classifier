package openai

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/swarnimcodes/website-productivity-classifier/utils"
)

func RetrieveAvailableModels(openaiBaseUrl, openAiApiKey string) (utils.Models, error) {
	endpoint := "v1/models"
	url := fmt.Sprintf("%s/%s", openaiBaseUrl, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return utils.Models{}, err
	}
	bearerToken := fmt.Sprintf("Bearer %s", openAiApiKey)
	req.Header.Set("Authorization", bearerToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return utils.Models{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		err := fmt.Sprintf("Non-2xx status code: %d", resp.StatusCode)
		return utils.Models{}, errors.New(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.Models{}, err
	}

	var models utils.Models
	err = json.Unmarshal([]byte(body), &models)
	if err != nil {
		return utils.Models{}, err
	}
	return models, nil
}
