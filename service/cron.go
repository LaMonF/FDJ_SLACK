package service

import (
	"github.com/LaMonF/FDJ_SLACK/model"
	"net/http"

	b "github.com/LaMonF/FDJ_SLACK/balance"
	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/parser"
	c "github.com/robfig/cron"
)

type cron struct{}

func SetUpCron() {
	cron := c.New()
	cron.AddFunc("0 15 22 * * MON,WED,SAT", func() { updateBalance(nil, nil) })
	cron.AddFunc("0 13 22 * * *", func() { Result(nil, nil) })
	cron.Start()
}

func updateBalance(w http.ResponseWriter, r *http.Request) {
	p := parser.NewParser()
	result, err := p.GetLotteryResult()
	if err != nil {
		l.Err("UpdateBalance", err)
	} else {
		b.CurrentBalance.UpdateBalance(result, model.MyBet)
		PostToSlack(b.CurrentBalance.String(), w)
	}
}
