package main

import (
	l "FDJ_SLACK/log"
	"FDJ_SLACK/parser"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var win = []int{7, 14, 22, 28, 42}
var luckWin = 5

var slackURL = os.Getenv("SLACK_HOOK_URL")
var localhost = "http://localhost:8888"

func main() {
	if slackURL == "" {
		slackURL = localhost
	}

	p := parser.NewParser()
	data := p.FetchData()
	result := p.ParseData(data)
	if len(result) > 0 {
		r := result[0]
		l.Info(r)
		postToSlack(slackURL, r.String())
		if r.IsWinning(win, luckWin) {
			postToSlack(slackURL, "ON A GAGNÉ !!!")
		}
	}
}

func postToSlack(slackURL string, post string) {
	message := fmt.Sprintf(`{"text": "%s"}`, post)

	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer([]byte(message)))
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
