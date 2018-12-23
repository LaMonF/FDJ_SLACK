package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	l "github.com/LaMonF/FDJ_SLACK/log"
	"github.com/LaMonF/FDJ_SLACK/model"
	"github.com/LaMonF/FDJ_SLACK/utils"

	"github.com/beevik/etree"
)

var apiURL = "https://www.lesbonsnumeros.com/loto/rss.xml"

type LesBonsNumerosAPIParser struct {
}

func NewParser() *LesBonsNumerosAPIParser {
	s := &LesBonsNumerosAPIParser{}
	return s
}

func (p *LesBonsNumerosAPIParser) GetLotteryResult() (model.LotteryResult, error) {
	var lastResult model.LotteryResult

	data, err := p.fetchData()
	if err != nil {
		return lastResult, err
	}

	results, err := p.parseData(data)
	if err != nil {
		return lastResult, err
	}

	for index, result := range results {
		if index == 0 { // only first result
			l.Info(result)
			//We can improve this post by using the URL from the POST request
			//See (https://api.slack.com/slash-commands -> Sending delayed responses)
			return result, nil
		}
	}
	return lastResult, errors.New("Last Result not found")
}

func (p *LesBonsNumerosAPIParser) fetchData() ([]byte, error) {
	l.Info("Get data from ApiURL ", apiURL)

	//Get data from URL
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("can't fetch data: %v", err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read fetched data: %v", err)
	}
	return contents, nil
}

func (p *LesBonsNumerosAPIParser) parseData(data []byte) ([]model.LotteryResult, error) {
	l.Info("Parsing data")

	//Create new XML Document to go through results
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	var result []model.LotteryResult

	//For each item (using XPath to get description)
	for _, t := range doc.FindElements("//item/description") {
		l.Info("Found item")
		for _, child := range t.Child {
			if c, ok := child.(*etree.CharData); ok {

				// Manage only not empty strings
				if !utils.IsWhitespace(c.Data) {

					//Cleaning Result to remove any HTML tags (ex: <h1>)
					cleanHTML := utils.CleanHTML(c.Data)

					//Get result line by line
					resultByLine := strings.Split(cleanHTML, "\n")

					l.Info(resultByLine)

					date := extractDateLottery(resultByLine[1])
					balls, err := extractBalls(resultByLine[2])
					if err != nil {
						return nil, err
					}

					luckyBall, err := extractLuckyBall(resultByLine[3])
					if err != nil {
						return nil, err
					}
					numberWinner := 0
					winnerPrize := 0
					nextLotteryDate := "Unknown"
					nextLotteryPrize := 0
					if len(resultByLine) > 10 { // When results are not up to date len(result_line_list) < 10
						numberWinner = extractNumberWinner(resultByLine[11])
						if winnerPrize, err = extractWinnerPrize(resultByLine[11]); err != nil {
							return nil, err
						}
						nextLotteryDate = extractNextLotteryDate(resultByLine[len(resultByLine)-2])
						if nextLotteryPrize, err = extractNextLotteryPrize(resultByLine[len(resultByLine)-2]); err != nil {
							return nil, err
						}
					}

					lotteryResult := model.LotteryResult{
						Date:             date,
						Balls:            balls,
						LuckyBall:        luckyBall,
						WinnerNumber:     numberWinner,
						WinnerPrize:      winnerPrize,
						NextLotteryDate:  nextLotteryDate,
						NextLotteryPrize: nextLotteryPrize,
					}

					result = append(result, lotteryResult)
				}
			}
		}
	}

	return result, nil
}

func extractDateLottery(line string) string {
	return line[len("Résultats du "):]
}

func extractBalls(line string) ([]int, error) {
	line = line[len("Numéros : "):]
	ballsAsString := strings.Split(line, " - ")

	var balls []int
	for _, ballAsString := range ballsAsString {
		i, err := strconv.Atoi(ballAsString)
		if err != nil {
			return nil, fmt.Errorf("can't extract balls: %v", err)
		}
		balls = append(balls, i)
	}
	return balls, nil
}

func extractLuckyBall(line string) (int, error) {
	luckyBallAsString := line[len("Numéro Chance : "):]
	i, err := strconv.Atoi(luckyBallAsString)
	if err != nil {
		return 0, fmt.Errorf("can't extract lucky ball: %v", err)
	}
	return i, nil
}

func extractNumberWinner(line string) int {
	winnerNumber := -1
	line = strings.ToLower(line)
	if strings.Contains(line, "le jackpot n'a pas été remporté lors de ce tirage !") {
		winnerNumber = 0
	} else if strings.Contains(line, utils.ONE) {
		winnerNumber = 1
	} else if strings.Contains(line, utils.TWO) {
		winnerNumber = 2
	} else if strings.Contains(line, utils.THREE) {
		winnerNumber = 3
	} else if strings.Contains(line, utils.FOUR) {
		winnerNumber = 4
	} else if strings.Contains(line, utils.FIVE) {
		winnerNumber = 5
	} else if strings.Contains(line, utils.SIX) {
		winnerNumber = 6
	} else if strings.Contains(line, utils.SEVEN) {
		winnerNumber = 7
	} else if strings.Contains(line, utils.EIGHT) {
		winnerNumber = 8
	} else if strings.Contains(line, utils.NINE) {
		winnerNumber = 9
	} else if strings.Contains(line, utils.TEN) {
		winnerNumber = 10
	}
	return winnerNumber
}

// Could be done in a better way
func extractWinnerPrize(line string) (int, error) {
	indexBegin := strings.Index(line, "montant de") + len("montant de")
	if indexBegin <= len("montant de") {
		return 0, nil
	}
	indexEnd := strings.Index(line, "€")
	prizeStr := strings.Replace(line[indexBegin:indexEnd], " ", "", -1)
	prizeStr = strings.Replace(prizeStr, "&nbsp;", "", -1)
	prize, err := strconv.Atoi(prizeStr)
	if err != nil {
		return 0, fmt.Errorf("can't return winner prize: %v", err)
	}
	return prize, nil
}

func extractNextLotteryDate(line string) string {
	line = line[len("Le montant du jackpot du prochain tirage du "):]
	index := strings.Index(line, "est")
	return line[:index]
}

// Could be done in a better way
func extractNextLotteryPrize(line string) (int, error) {
	indexBegin := strings.Index(line, "est de") + len("est de")
	indexEnd := strings.Index(line, "€")
	prizeStr := strings.Replace(line[indexBegin:indexEnd], " ", "", -1)
	prize, err := strconv.Atoi(prizeStr)
	if err != nil {
		return 0, fmt.Errorf("can't extract next lottery prize: %v", err)
	}
	return prize, nil
}
