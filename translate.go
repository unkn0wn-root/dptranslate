package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Translation struct {
	TranslatedText string `json:"translatedText"`
	SourceLanguage string `json:"sourceLanguage"`
}

type Translate struct {
	APIKey string
	To     string
	Text   string
	ProAPI bool
}

// init a translate struct with default values
func NewTranslate() *Translate {
	return &Translate{}
}

// generate url scheme based on ProAPI flag
func (t *Translate) generateURLScheme() string {
	if t.ProAPI {
		return "https://api.deepl.com"
	}

	return "https://api-free.deepl.com"
}

func (t *Translate) Translate() (*Translation, error) {
	urlScheme := t.generateURLScheme()
	payload := map[string]string{
		"auth_key":    t.APIKey,
		"text":        t.Text,
		"target_lang": t.To,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlScheme+"/v2/translate", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// DeepL returns 200. Assume error if != 200
	if resp.StatusCode != http.StatusOK {
		errResponse, err := checkErrorResponseBody(resp.Body)
		if err != nil {
			log.Print(err)
		}

		apiError := NewAPIError(resp.StatusCode, errResponse)
		return nil, apiError
	}

	var translation Translation
	if err := json.NewDecoder(resp.Body).Decode(&translation); err != nil {
		return nil, err
	}

	return &translation, nil
}

// check for error response in body (if any)
func checkErrorResponseBody(body io.Reader) (string, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	errorMessage := func() string {
		if len(bodyBytes) == 0 {
			return "Could not make request to DeepL API"
		}

		return string(bodyBytes)
	}()

	return errorMessage, nil
}
