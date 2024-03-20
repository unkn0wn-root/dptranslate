package deepl

import (
	"flag"
)

func ParseFlags(t *Translate) {
	flag.StringVar(&t.APIKey, "api-key", "", "DeepL API Key. If you do not want to provide Api key here - use env variable DEEPL_API_KEY")
	flag.StringVar(&t.To, "to", "", "Translate to language")
	flag.StringVar(&t.Text, "text", "", "Source text")
	flag.BoolVar(&t.ProAPI, "pro-api", false, "Use free DeepL API (default: false)")

	flag.Parse()
}
