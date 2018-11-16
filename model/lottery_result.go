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

func (l *LotteryResult) IsWinning(listWinningBalls []int, winningLuckyBall int) bool {
	sort.Ints(listWinningBalls)
	return utils.TestEq(listWinningBalls, l.Balls) && winningLuckyBall == l.LuckyBall
}

func (l *LotteryResult) String() string {
	var sb strings.Builder

	sb.WriteString("Résultats du ")
	sb.WriteString(l.Date)
	sb.WriteString("\n")

	sb.WriteString("Numéros: ")
	for _, ball := range l.Balls {
		sb.WriteString(strconv.Itoa(ball))
		sb.WriteString(" ")
	}
	sb.WriteString("\n")

	sb.WriteString("Numéro chance : ")
	sb.WriteString(strconv.Itoa(l.LuckyBall))
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