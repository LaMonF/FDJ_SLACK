package balance

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/model"
)

func TestBalance_getBalanceValue(t *testing.T) {
	// Init
	dir, err := ioutil.TempDir("", "fdjSlackTEST-")
	if err != nil {
		l.Error(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "TestBalance_getBalanceValue")

	const EXPECTED_VALUE = 666

	content := []byte(strconv.Itoa(EXPECTED_VALUE))
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		l.Error(err)
	}

	result := getBalanceValue(tmpfn)
	if result != EXPECTED_VALUE {
		t.Fail()
	}
}

func TestBalance_SetBalanceValue(t *testing.T) {
	// Init
	dir, err := ioutil.TempDir("", "fdjSlackTEST-")
	if err != nil {
		l.Error(err)
	}
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "TestBalance_SetBalanceValue")

	const EXPECTED_VALUE = 999
	balance := NewBalance(tmpfn)
	balance.SetBalanceValue(EXPECTED_VALUE)
	if balance.Value != EXPECTED_VALUE {
		t.Fail()
	}

	dat, err := ioutil.ReadFile(tmpfn)
	formattedString := strings.Replace(string(dat), "\n", "", -1)
	value, _ := strconv.ParseFloat(formattedString, 64)
	if value != EXPECTED_VALUE {
		t.Fail()
	}
}

func TestBalance_CalculateBalance_RANKING2(t *testing.T) {
	// Init
	dir, err := ioutil.TempDir("", "fdjSlackTEST-")
	if err != nil {
		l.Error(err)
	}
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "TestBalance_CalculateBalance_RANKING2.fdjSlack")

	balance := NewBalance(tmpfn)
	balance.Value = 10

	// RANK 2
	lotteryResult := model.LotteryResult{
		Date:             "",
		Balls:            []int{7, 14, 22, 28, 42},
		LuckyBall:        9,
		WinnerNumber:     1,
		WinnerPrize:      10000000,
		NextLotteryDate:  "",
		NextLotteryPrize: 1000000,
	}

	bet := model.BetCombo{
		Balls: []int{7, 14, 22, 28, 42},
		Bonus: 5,
	}

	balance.CalculateBalance(lotteryResult, bet)

	const EXPECTED_VALUE = 100007.80
	dat, err := ioutil.ReadFile(tmpfn)
	formattedString := strings.Replace(string(dat), "\n", "", -1)
	value, _ := strconv.ParseFloat(formattedString, 64)
	if value != EXPECTED_VALUE {
		t.Fail()
	}
}

func TestBalance_CalculateBalance_RANKING0(t *testing.T) {
	// Init
	dir, err := ioutil.TempDir("", "fdjSlackTEST-")
	if err != nil {
		l.Error(err)
	}
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "TestBalance_CalculateBalance_RANKING0.fdjSlack")

	balance := NewBalance(tmpfn)
	balance.Value = 10

	// RANK 0
	lotteryResult := model.LotteryResult{
		Date:             "",
		Balls:            []int{1, 2, 3, 4, 5},
		LuckyBall:        9,
		WinnerNumber:     1,
		WinnerPrize:      10000000,
		NextLotteryDate:  "",
		NextLotteryPrize: 1000000,
	}

	bet := model.BetCombo{
		Balls: []int{7, 14, 22, 28, 42},
		Bonus: 5,
	}

	balance.CalculateBalance(lotteryResult, bet)

	const EXPECTED_VALUE = 7.80
	dat, err := ioutil.ReadFile(tmpfn)
	formattedString := strings.Replace(string(dat), "\n", "", -1)
	value, _ := strconv.ParseFloat(formattedString, 64)
	if value != EXPECTED_VALUE {
		t.Fail()
	}
}

func TestBalance_CalculateBalance_RANKING9(t *testing.T) {
	// Init
	dir, err := ioutil.TempDir("", "fdjSlackTEST-")
	if err != nil {
		l.Error(err)
	}
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "TestBalance_CalculateBalance_RANKING9.fdjSlack")

	balance := NewBalance(tmpfn)
	balance.Value = 10

	// RANK 0
	lotteryResult := model.LotteryResult{
		Date:             "",
		Balls:            []int{1, 2, 3, 4, 5},
		LuckyBall:        5,
		WinnerNumber:     1,
		WinnerPrize:      10000000,
		NextLotteryDate:  "",
		NextLotteryPrize: 1000000,
	}

	bet := model.BetCombo{
		Balls: []int{7, 14, 22, 28, 42},
		Bonus: 5,
	}

	balance.CalculateBalance(lotteryResult, bet)

	const EXPECTED_VALUE = 10
	dat, err := ioutil.ReadFile(tmpfn)
	formattedString := strings.Replace(string(dat), "\n", "", -1)
	value, _ := strconv.ParseFloat(formattedString, 64)
	if value != EXPECTED_VALUE {
		t.Fail()
	}
}

func TestBalance_CalculateBalance_RANKING_NOMONEY(t *testing.T) {
	// Init
	dir, err := ioutil.TempDir("", "fdjSlackTEST-")
	if err != nil {
		l.Error(err)
	}
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "TestBalance_CalculateBalance_RANKING_NOMONEY.fdjSlack")

	const EXPECTED_VALUE = 10

	content := []byte(strconv.Itoa(EXPECTED_VALUE))
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		l.Error(err)
	}

	balance := NewBalance(tmpfn)
	balance.Value = 1

	// RANK 0
	lotteryResult := model.LotteryResult{
		Date:             "",
		Balls:            []int{1, 2, 3, 4, 5},
		LuckyBall:        5,
		WinnerNumber:     1,
		WinnerPrize:      10000000,
		NextLotteryDate:  "",
		NextLotteryPrize: 1000000,
	}

	bet := model.BetCombo{
		Balls: []int{7, 14, 22, 28, 42},
		Bonus: 5,
	}

	balance.CalculateBalance(lotteryResult, bet)

	dat, err := ioutil.ReadFile(tmpfn)
	formattedString := strings.Replace(string(dat), "\n", "", -1)
	value, _ := strconv.ParseFloat(formattedString, 64)
	if value != EXPECTED_VALUE {
		t.Fail()
	}
}
