package cli

import (
	"github.com/Jeffail/gabs"
	"log"
	"net/http"
	"sync"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslation(reqBody *RequestBody, str chan string, wg *sync.WaitGroup) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", translateUrl, nil)
	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", reqBody.SourceLang)
	query.Add("tl", reqBody.TargetLang)
	query.Add("dt", "t")
	query.Add("q", reqBody.SourceText)
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "Too many requests"
		wg.Done()
		return
	}
	processedResponse, err := gabs.ParseJSONBuffer(res.Body)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	nestOne, err := processedResponse.ArrayElement(0)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	translatedText, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	str <- translatedText.Data().(string)
	wg.Done()
}
