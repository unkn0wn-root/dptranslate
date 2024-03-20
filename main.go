package main

import (
	deepl "deepl/translate"
	"fmt"
	"log"
	"os"
)

func main() {
	translate := deepl.NewTranslate()
	deepl.ParseFlags(translate)

	// check for api key from args. Fallback to env var if key is not provided by the user
	if translate.APIKey == "" {
		apiKey := os.Getenv("DEEPL_API_KEY")
		if apiKey == "" {
			log.Fatal("No Deepl Api found! Please set the DEEPL_API_KEY environment variable or use flag --api-key")
		}
	}

	translation, err := translate.Translate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Translated Text: %s\nSource Language: %s\n", translation.TranslatedText, translation.SourceLanguage)
}
