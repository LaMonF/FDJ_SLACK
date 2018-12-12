package model

import (
	"strconv"
	"strings"
)

var MyBet = BetCombo{
	Balls: []int{7, 14, 22, 28, 42},
	Bonus: 5,
}

type BetCombo struct {
	Balls []int
	Bonus int
}

func (b *BetCombo) String() string {
	var sb strings.Builder

	sb.WriteString("Numéros: ")
	for _, ball := range b.Balls {
		sb.WriteString(strconv.Itoa(ball))
		sb.WriteString(" ")
	}
	sb.WriteString("\n")

	sb.WriteString("Numéro chance : ")
	sb.WriteString(strconv.Itoa(b.Bonus))
	sb.WriteString("\n")

	return sb.String()
}
