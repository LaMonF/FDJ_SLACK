package main

import (
	l "FDJ_SLACK/log"
	"FDJ_SLACK/parser"
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

var win = []int{7, 14, 22, 28, 42}
var luckWin = 5

var slackURL = os.Getenv("SLACK_HOOK_URL")
var localhost = "http://localhost:8888"
var apiURL = "https://www.lesbonsnumeros.com/loto/rss.xml"

func main() {
	if slackURL == "" {
		slackURL = localhost
	}

	p := parser.NewParser()
	result := p.GetAndParseData(apiURL)
	for index, result := range result {
		if index == 0 { // only first result
			l.Info(result)
			postToSlack(slackURL, result.String())
			if result.IsWinning(win, luckWin) {
				postToSlack(slackURL, "ON A GAGNÉ !!!")
			}
		}
	}
}

func postToSlack(slackURL string, post string) {
	message := `{"text" : "` + post + `"}`

	var jsonStr = []byte(message)
	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	//TODO: Manage response, handle errors
	l.Info("response Body:", string(body))
}
