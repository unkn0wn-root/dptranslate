package main

import (
	"flag"
	"fmt"
	"os"
)

func ParseFlags(t *Translate) {
	flag.StringVar(&t.APIKey, "api-key", "", "DeepL API Key. If you do not want to provide Api key here - use env variable DEEPL_API_KEY")
	flag.StringVar(&t.APIKey, "a", "", "DeepL API Key. If you do not want to provide Api key here - use env variable DEEPL_API_KEY (shorthand)")
	flag.StringVar(&t.To, "to", "", "Translate to language")
	flag.StringVar(&t.To, "t", "", "Translate to language (shorthand)")
	flag.StringVar(&t.Text, "text", "", "Source text")
	flag.StringVar(&t.Text, "x", "", "Source text (shorthand)")
	flag.BoolVar(&t.ProAPI, "pro-api", false, "Use free DeepL API (default: false)")
	flag.BoolVar(&t.ProAPI, "p", false, "Use free DeepL API (default: false) (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, " -a --api-key \t DeepL API key. Use env. var if you dont want to provide it here\n")
		fmt.Fprintf(os.Stderr, " -t --to \t Translate given text to target language\n")
		fmt.Fprintf(os.Stderr, " -x --text \t Source text to translate\n")
		fmt.Fprintf(os.Stderr, " -p --pro-api \t Use free DeepL API (default: false)\n")
	}

	flag.Parse()
}
