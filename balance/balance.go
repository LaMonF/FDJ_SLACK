package balance

import (
	"bufio"
	"fmt"
	l "github.com/LaMonF/FDJ_SLACK/log"
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



