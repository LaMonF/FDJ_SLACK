package model

import (
	"strconv"
	"strings"
)

type BetCombo struct {
	Balls    []int
	Bonus    int
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

