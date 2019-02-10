package balance

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/LaMonF/FDJ_SLACK/conf"
	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/model"
)

type Balance struct {
	filename string
	Value    float64
}

func NewBalance(fileName string) Balance {
	balance := Balance{}
	balance.filename = fileName
	balance.Value = getBalanceValue(fileName)
	return balance
}

var CurrentBalance = NewBalance(conf.BALANCE_FILE_PATH)

// PUBLIC

func (b *Balance) SetBalanceValue(value float64) {
	b.Value = value
	writeFile(b.filename, value)
}

func (b *Balance) CalculateBalance(result model.LotteryResult, bet model.BetCombo) {
	if b.Value > conf.BET_PRICE {
		b.Value = b.Value - conf.BET_PRICE
		winRankingBalance := GetwinRanking(result, bet)
		b.Value = b.Value + float64(winRankingBalance)
		writeFile(b.filename, b.Value)
		l.Debug("New balance : " + b.String())
	} else {
		l.Error("Not enough money left : " + b.String())
	}
}

func (b *Balance) String() string {
	var sb strings.Builder
	sb.WriteString("Solde courant : ")
	sb.WriteString(strconv.FormatFloat(b.Value, 'f', 2, 64))
	sb.WriteString(" € \n")
	return sb.String()
}

func (b *Balance) StringWinning(l model.LotteryResult, bet model.BetCombo) string {
	var sb strings.Builder
	var winning = float64(GetwinRanking(l, bet)) - conf.BET_PRICE
	sb.WriteString("Gains : ")
	sb.WriteString(strconv.FormatFloat(winning, 'f', 2, 64))
	sb.WriteString(" € \n")
	return sb.String()
}

// PRIVATE

func getBalanceValue(fileName string) float64 {
	file, err := os.Open(fileName)
	if err != nil {
		l.Error("Cannot Open file", fileName, ":", err)
		return 0
	}
	defer file.Close()
	return readFile(file)
}

// File

func readFile(file *os.File) float64 {
	dat, err := ioutil.ReadAll(file)
	if err != nil {
		l.Error("readFile", err)
	}
	formattedString := strings.Replace(string(dat), "\n", "", -1)
	value, _ := strconv.ParseFloat(formattedString, 64)
	return value
}

func writeFile(fileName string, value float64) {
	d1 := []byte(fmt.Sprintf("%.2f", value))
	ioutil.WriteFile(fileName, d1, 0644)
}
