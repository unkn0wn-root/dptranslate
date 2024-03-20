package deepl

import (
	"bytes"
	"encoding/json"
	"errors"
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
func (t *Translate) GenerateURLSceheme() string {
	if t.ProAPI {
		return "https://api.deepl.com"
	}

	return "https://api-free.deepl.com"
}

// perform the actual translation
func (t *Translate) Translate() (*Translation, error) {
	urlScheme := t.GenerateURLSceheme()
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

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		apiError := APIError{StatusCode: resp.StatusCode, Body: string(bodyBytes)}.Error()
		return nil, errors.New(apiError)
	}

	var translation Translation
	if err := json.NewDecoder(resp.Body).Decode(&translation); err != nil {
		return nil, err
	}

	return &translation, nil
}
