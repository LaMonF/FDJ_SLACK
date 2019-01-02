package balance

import (
	"github.com/LaMonF/FDJ_SLACK/model"
	"github.com/LaMonF/FDJ_SLACK/utils"
	"strings"
)

type Paytable struct{}

// It defines the winning rank which determines the amount of winning money.
// cf : https://github.com/LaMonF/FDJ_SLACK/issues/10
type WIN_RANK float64

const (
	RANK_1 WIN_RANK = 2000000
	RANK_2 WIN_RANK = 100000
	RANK_3 WIN_RANK = 1000
	RANK_4 WIN_RANK = 500
	RANK_5 WIN_RANK = 50
	RANK_6 WIN_RANK = 20
	RANK_7 WIN_RANK = 10
	RANK_8 WIN_RANK = 5
	RANK_9 WIN_RANK = 2.20
	RANK_0 WIN_RANK = 0
)

func GetwinRanking(result model.LotteryResult, bet model.BetCombo) WIN_RANK {
	var occurence = utils.ArrayNumberSameOccurence(result.Balls, bet.Balls)
	if result.LuckyBall == bet.Bonus {
		if occurence == 0 {
			return RANK_9
		}
		if occurence == 1 {
			return RANK_9
		}
		if occurence == 2 {
			return RANK_7
		}
		if occurence == 3 {
			return RANK_5
		}
		if occurence == 4 {
			return RANK_3
		}
		if occurence == 5 {
			return RANK_1
		}
	} else {
		if occurence == 2 {
			return RANK_8
		}
		if occurence == 3 {
			return RANK_6
		}
		if occurence == 4 {
			return RANK_4
		}
		if occurence == 5 {
			return RANK_2
		}
	}
	return RANK_0
}

func PaytableString() string {
	var sb strings.Builder
	sb.WriteString("5 bons numéros + numéro chance   --> JACKPOT \n")
	sb.WriteString("5 bons numéros                   --> 100000€ \n")
	sb.WriteString("4 bons numéros + numéro chance   --> 1000€   \n")
	sb.WriteString("4 bons numéros                   --> 500€    \n")
	sb.WriteString("3 bons numéros + numéro chance   --> 50€     \n")
	sb.WriteString("3 bons numéros                   --> 20€     \n")
	sb.WriteString("2 bons numéros + numéro chance   --> 10€     \n")
	sb.WriteString("2 bons numéros                   --> 5€      \n")
	sb.WriteString("1 bon  numéro  + numéro chance   --> 2.20€   \n")
	sb.WriteString("0 bon numéro   + numéro chance   --> 2.20€   \n")
	return sb.String()
}
