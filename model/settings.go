package model

import (
	"fmt"
	"strings"
)

type Settings struct {
	BalanceFile   string   `yaml:"balance_file"`
	BetPrice      float64  `yaml:"bet_price"`
	Bet           BetCombo `yaml:"bet"`
	CronPostSlack string   `yaml:"cron_post_slack"`
}

func (s *Settings) String() string {
	var sb strings.Builder
	sb.WriteString("Ficher solde courant: " + s.BalanceFile)
	sb.WriteString("\n")
	sb.WriteString("Prix de la grille: " + fmt.Sprintf("%.2f", s.BetPrice))
	sb.WriteString("\n")
	sb.WriteString(s.Bet.String())
	sb.WriteString("\n")
	sb.WriteString("CRONTAB pour post sur slack: " + s.CronPostSlack)
	return sb.String()
}