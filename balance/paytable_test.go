package balance

import (
	"github.com/LaMonF/FDJ_SLACK/model"
	"testing"
)

func TestPaytable_GetwinRanking_RANK2(t *testing.T) {

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

	result := GetwinRanking(lotteryResult, bet)

	if result != RANK_2 {
		t.Fail()
	}
}

func TestPaytable_GetwinRanking_RANK0(t *testing.T) {

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

	result := GetwinRanking(lotteryResult, bet)

	if result != RANK_0 {
		t.Fail()
	}
}

func TestPaytable_GetwinRanking_RANK9(t *testing.T) {

	// RANK 9
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

	result := GetwinRanking(lotteryResult, bet)

	if result != RANK_9 {
		t.Fail()
	}
}
