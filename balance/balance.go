package balance

import (
	"bufio"
	"fmt"
	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/model"
	"github.com/LaMonF/FDJ_SLACK/utils"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const BALANCE_FILE_PATH  = "/tmp/balance.fdjSlack"


type Balance struct {
	Value    	float64
	File 		*os.File
}

func NewBalance() Balance {
	balance := Balance{}

	file, err := os.Open(BALANCE_FILE_PATH)
	if err != nil {
		os.Create(BALANCE_FILE_PATH)
	}
	balance.File = file
	balance.Value = balance.readFile()
	balance.File.Close()
	return balance
}


func (b *Balance) readFile() float64 {
	dat, err := ioutil.ReadAll(b.File)
	if err != nil {
		l.Error("readFile", err)
	}
	formattedString := strings.Replace(string(dat), "\n", "", -1)
	value,_ := strconv.ParseFloat(formattedString, 64);
	return value;
}

func (b *Balance) writeFile(value float64) {
	// Create a buffered writer from the file
	bufferedWriter := bufio.NewWriter(b.File)
	fmt.Fprint(bufferedWriter,"%.2f", b.Value)
	b.File.Close()
}

func (b *Balance) String() string{
	var sb strings.Builder
	sb.WriteString("Solde courant : ")
	sb.WriteString(strconv.FormatFloat(b.Value, 'f', 2, 64))
	sb.WriteString(" € \n")
	return sb.String()
}

func (b *Balance) updateBalance(result model.LotteryResult, bet model.BetCombo) {
	if b.Value > 2.20 {
		b.Value = b.Value - 2.20 // Price of a bet
		winRankingBalance := getwinRanking(result, bet)
		b.Value = b.Value + float64(winRankingBalance)
		b.writeFile(b.Value)
		l.Debug("New balance : "+ b.String())
	} else {
		l.Error("Not enough money left : " + b.String())
	}
}

func getwinRanking(result model.LotteryResult, bet model.BetCombo) utils.WIN_RANK {
	var occurence= utils.ArrayNumberSameOccurence(result.Balls, bet.Balls)
	if result.LuckyBall == bet.Bonus {
		if occurence == 0 {
			return utils.RANK_9
		}
		if occurence == 1 {
			return utils.RANK_9
		}
		if occurence == 2 {
			return utils.RANK_7
		}
		if occurence == 3 {
			return utils.RANK_5
		}
		if occurence == 4 {
			return utils.RANK_3
		}
		if occurence == 5 {
			return utils.RANK_1
		}
	} else {
		if (occurence == 2) {
			return utils.RANK_8
		}
		if (occurence == 3) {
			return utils.RANK_6
		}
		if (occurence == 4) {
			return utils.RANK_4
		}
		if (occurence == 5) {
			return utils.RANK_2
		}
	}
	return utils.RANK_0
}




