package service

import (
	"fmt"
	"net/http"
	"strconv"

	b "github.com/LaMonF/FDJ_SLACK/balance"
	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/model"
	"github.com/LaMonF/FDJ_SLACK/utils"
)

var slackURL = utils.GetEnv("SLACK_HOOK_URL", "http://localhost:8888")

const API_VERSION = 1

const (
	LOTORESULT string = "lotoResult";
	BETBALLS   string = "balls";
	BALANCE    string = "balance";
)

type api struct{}

func SetUpServer() {
	http.HandleFunc(fmt.Sprintf("/%d/%s", API_VERSION, LOTORESULT), Result)
	http.HandleFunc(fmt.Sprintf("/%d/%s", API_VERSION, BETBALLS), betBalls)
	http.HandleFunc(fmt.Sprintf("/%d/%s", API_VERSION, BALANCE), balance)
	// set router
	err := http.ListenAndServe(":9090", nil)
	// set listen port
	if err != nil {
		l.Err("ListenAndServe: ", err)
	}
}

func betBalls(w http.ResponseWriter, r *http.Request) {
	PostToSlack(model.MyBet.String(), w)
}

func balance(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newBalance, err := strconv.ParseFloat(r.FormValue("text"), 64);
	if err != nil {
		l.Info("Fail to read the new balance as a Float64")
	} else {
		b.CurrentBalance.WriteFile(newBalance)
	}

	PostToSlack(b.CurrentBalance.String(), w)
}
