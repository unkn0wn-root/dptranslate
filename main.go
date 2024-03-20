package main

import (
	deepl "deepl/translate"
	"fmt"
	"os"
)

func main() {
	translate := deepl.NewTranslate()
	deepl.ParseFlags(translate)

	if translate.APIKey == "" {
		apiKey := os.Getenv("DEEPL_API_KEY")
		if apiKey == "" {
			fmt.Println("No Deepl Api found! Please set the DEEPL_API_KEY environment variable or use flag --api-key")
			os.Exit(1)
		}
	}

	translation, err := translate.Translate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Translated Text: %s\nSource Language: %s\n", translation.TranslatedText, translation.SourceLanguage)
}
