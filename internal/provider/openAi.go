package provider

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type OpenAICompletionsRequest struct {
	Model    Model     `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenAICompletionsResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func CallOpenAICompatibleCompletionsAPI(url, key string, request OpenAICompletionsRequest) (OpenAICompletionsResponse, error) {
	var chatResponse OpenAICompletionsResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return chatResponse, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return chatResponse, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	Logger.Printf("Sending request to OpenAI-compatible API. URL: %s, Request: %s", url, jsonReq)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return chatResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return chatResponse, err
	}

	Logger.Printf("Received response from OpenAI-compatible API. Response: %s", body)

	if err = json.Unmarshal(body, &chatResponse); err != nil {
		return chatResponse, err
	}

	return chatResponse, nil
}
