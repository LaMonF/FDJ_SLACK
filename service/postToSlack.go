package service

import (
	"bytes"
	"github.com/LaMonF/FDJ_SLACK/conf"
	"net/http"
	"strings"

	b "github.com/LaMonF/FDJ_SLACK/balance"
	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/parser"
)

func Result(w http.ResponseWriter, r *http.Request) {
	p := parser.NewParser()
	result, err := p.GetLotteryResult()
	if err != nil {
		l.Err("getResultAndPostToSlack", err)
	} else {
		var builder strings.Builder
		builder.WriteString(result.String(conf.Settings.Bet))
		builder.WriteString(b.CurrentBalance.StringWinning(result, conf.Settings.Bet))
		PostToSlack(builder.String(), w)
		if result.IsWinning(conf.Settings.Bet) {
			PostToSlack("ON A GAGNÉ !!!", w)
		}
	}
}

func PostToSlack(post string, w http.ResponseWriter) {
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

	l.Info("Sending back response to Slack : POST --> " + message)

	var jsonStr = []byte(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(jsonStr)
}
