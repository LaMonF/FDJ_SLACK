package model

import (
	"sort"
	"strconv"
	"strings"

	"github.com/LaMonF/FDJ_SLACK/utils"
)

type LotteryResult struct {
	Date             string
	Balls            []int
	LuckyBall        int
	WinnerNumber     int
	WinnerPrize      int
	NextLotteryDate  string
	NextLotteryPrize int
}

func (l *LotteryResult) IsWinning(bet BetCombo) bool {
	sort.Ints(bet.Balls)
	return utils.TestEq(bet.Balls, l.Balls) && bet.Bonus == l.LuckyBall
}

func (l *LotteryResult) String(bet BetCombo) string {
	var sb strings.Builder

	sb.WriteString("Résultats du ")
	sb.WriteString(l.Date)
	sb.WriteString("\n")

	sb.WriteString("Numéros: ")
	for _, ball := range l.Balls {
		if utils.Contains(bet.Balls, ball) { sb.WriteString("*")}
		sb.WriteString(strconv.Itoa(ball))
		if utils.Contains(bet.Balls, ball) { sb.WriteString("*")}
		sb.WriteString(" ")
	}
	sb.WriteString("\n")

	sb.WriteString("Numéro chance : ")
	if l.LuckyBall == bet.Bonus { sb.WriteString("*")}
	sb.WriteString(strconv.Itoa(l.LuckyBall))
	if l.LuckyBall == bet.Bonus { sb.WriteString("*")}
	sb.WriteString("\n")

	sb.WriteString(l.GetCurrentWinnerString())
	sb.WriteString("\n")

	sb.WriteString("Le prochain tirage sera le ")
	sb.WriteString(l.NextLotteryDate)
	sb.WriteString(" pour un montant de ")
	sb.WriteString(strconv.Itoa(l.NextLotteryPrize))
	sb.WriteString(" €.\n")

	return sb.String()
}

func (l *LotteryResult) GetCurrentWinnerString() string {
	if l.WinnerNumber == 1 {
		var sb strings.Builder
		sb.WriteString("Un joueur a remporté le jackpot d'un montant de ")
		sb.WriteString(strconv.Itoa(l.WinnerPrize))
		sb.WriteString(" €.")
		return sb.String()
	} else if l.WinnerNumber > 1 {
		var sb strings.Builder
		sb.WriteString("Le jackpot a été remporté par ")
		sb.WriteString(strconv.Itoa(l.WinnerNumber))
		sb.WriteString(" joueurs, ils se partagent ")
		sb.WriteString(strconv.Itoa(l.WinnerPrize))
		sb.WriteString(" €.")
		return sb.String()
	}
	return "Le jackpot n'a pas été remporté lors de ce tirage !"
}
