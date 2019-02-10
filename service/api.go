package service

import (
	"fmt"
	"github.com/LaMonF/FDJ_SLACK/conf"
	"net/http"
	"strconv"

	b "github.com/LaMonF/FDJ_SLACK/balance"
	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/utils"
)

var slackURL = utils.GetEnv("SLACK_HOOK_URL", "http://localhost:8888")

const API_VERSION = 1

const (
	LOTORESULT string = "lotoResult";
	BETBALLS   string = "balls";
	BALANCE    string = "balance";
	PAYTABLE   string = "paytable"
)

type api struct{}

func SetUpServer() {
	http.HandleFunc(fmt.Sprintf("/%d/%s", API_VERSION, LOTORESULT), Result)
	http.HandleFunc(fmt.Sprintf("/%d/%s", API_VERSION, BETBALLS), betBalls)
	http.HandleFunc(fmt.Sprintf("/%d/%s", API_VERSION, BALANCE), balance)
	http.HandleFunc(fmt.Sprintf("/%d/%s", API_VERSION, PAYTABLE), paytable)
	// set router
	err := http.ListenAndServe(":9090", nil)
	// set listen port
	if err != nil {
		l.Err("ListenAndServe: ", err)
	}
}

func betBalls(w http.ResponseWriter, r *http.Request) {
	PostToSlack(conf.Settings.Bet.String(), w)
}

func balance(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newBalance, err := strconv.ParseFloat(r.FormValue("text"), 64);
	if err != nil {
		l.Info("Fail to read the new balance as a Float64")
	} else {
		b.CurrentBalance.SetBalanceValue(newBalance)
	}
	PostToSlack(b.CurrentBalance.String(), w)
}

func paytable(w http.ResponseWriter, r *http.Request) {
	PostToSlack(b.PaytableString(), w)
}
