package main

import (
	"bytes"
	"errors"
	"fmt"
	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/model"
	"github.com/LaMonF/FDJ_SLACK/parser"
	"github.com/LaMonF/FDJ_SLACK/utils"
	"github.com/robfig/cron"
	"net/http"
)


const (
	LOTORESULT string = "lotoResult"
)


const API_VERSION = 1

var win = []int{7, 14, 22, 28, 42}
var luckWin = 5

var slackURL = utils.GetEnv("SLACK_HOOK_URL", "http://localhost:8888")

func main() {
	startServer()
}

func startServer() {
	setUpCron()
	setUpServer()
}

func setUpServer() {
	http.HandleFunc(fmt.Sprintf("/%d/%s", API_VERSION, LOTORESULT), getResultAndPostToSlack)
	// set router
	err := http.ListenAndServe(":9090", nil)
	// set listen port
	if err != nil {
		l.Err("ListenAndServe: ", err)
	}
}

func setUpCron(){
	c := cron.New()
	c.AddFunc("0 0 21 * * *", func() { getResultAndPostToSlack(nil, nil) })
	c.Start()
}

func getResultAndPostToSlack(w http.ResponseWriter, r *http.Request) {
	result, err := getLotteryResult()
	if err != nil {
		l.Err("getResultAndPostToSlack", err)
	} else {
		postToSlack(result.String(), w)
		if result.IsWinning(win, luckWin) {
			postToSlack("ON A GAGNÉ !!!", w)
		}
	}
}

func getLotteryResult() (model.LotteryResult, error){
	p := parser.NewParser()
	data := p.FetchData()
	var lastResult model.LotteryResult
	results := p.ParseData(data)

	for index, result := range results {
		if index == 0 { // only first result
			l.Info(result)
			//We can improve this post by using the URL from the POST request
			//See (https://api.slack.com/slash-commands -> Sending delayed responses)
			return result, nil
		}
	}
	return lastResult, errors.New("Last Result not found")
}


func postToSlack(post string, w http.ResponseWriter) {
	if w != nil {
		sendResponseToSlack(w, post)
	} else {
		message := `{"text" : "` + post + `"}`
		var jsonStr = []byte(message)
		http.Post(slackURL, "application/json", bytes.NewBuffer(jsonStr))
	}
}

func sendResponseToSlack(w http.ResponseWriter, post string) {
	//We set the response_type to in_channel (everyone can see it) instead of ephemeral (only you) by default
	message := `{"response_type": "in_channel","text" : "` + post + `"}`

	l.Info("Sending back response to Slack : POST --> "+ message)

	var jsonStr = []byte(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(jsonStr)
}