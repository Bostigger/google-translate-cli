package main

import (
	"flag"
	"fmt"
	"github.com/bostigger/google-translate-cli/cli"
	"os"
	"strings"
	"sync"
)

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source Language[en]")
	flag.StringVar(&targetLang, "t", "fr", "Target Language")
	flag.StringVar(&sourceText, "st", "", "Text to translate")
}

var wg sync.WaitGroup

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	langChan := make(chan string)
	wg.Add(1)

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	go cli.RequestTranslation(reqBody, langChan, &wg)

	translatedResult := strings.ReplaceAll(<-langChan, "+", " ")
	fmt.Printf("%s\n", translatedResult)

	wg.Wait()
}
