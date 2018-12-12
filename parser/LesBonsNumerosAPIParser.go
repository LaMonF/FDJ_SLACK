package parser

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"errors"

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

	data := p.fetchData()
	results := p.parseData(data)

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

func (p *LesBonsNumerosAPIParser) fetchData() []byte {
	l.Info("Get data from ApiURL ", apiURL)

	//Get data from URL
	response, err := http.Get(apiURL)
	if err != nil {
		l.Error("%s", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		l.Error("%s", err)
		os.Exit(1)
	}
	return contents
}

func (p *LesBonsNumerosAPIParser) parseData(data []byte) []model.LotteryResult {
	l.Info("Parsing data")

	//Create new XML Document to go through results
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		panic(err)
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
					balls := extractBalls(resultByLine[2])
					luckyBall := extractLuckyBall(resultByLine[3])
					numberWinner := 0
					winnerPrize := 0
					nextLotteryDate := "Unknown"
					nextLotteryPrize := 0
					if len(resultByLine) > 10 { // When results are not up to date len(result_line_list) < 10
						numberWinner = extractNumberWinner(resultByLine[11])
						winnerPrize = extractWinnerPrize(resultByLine[11])
						nextLotteryDate = extractNextLotteryDate(resultByLine[len(resultByLine)-2])
						nextLotteryPrize = extractNextLotteryPrize(resultByLine[len(resultByLine)-2])
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

	return result
}

func extractDateLottery(line string) string {
	return line[len("Résultats du "):]
}

func extractBalls(line string) []int {
	line = line[len("Numéros : "):]
	ballsAsString := strings.Split(line, " - ")

	var balls []int
	for _, ballAsString := range ballsAsString {
		i, err := strconv.Atoi(ballAsString)
		if err != nil {
			// handle error
			l.Error(err)
			os.Exit(2)
		}
		balls = append(balls, i)
	}
	return balls
}

func extractLuckyBall(line string) int {
	luckyBallAsString := line[len("Numéro Chance : "):]
	i, err := strconv.Atoi(luckyBallAsString)
	if err != nil {
		// handle error
		l.Error(err)
		os.Exit(2)
	}
	return i
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
func extractWinnerPrize(line string) int {
	indexBegin := strings.Index(line, "montant de") + len("montant de")
	if indexBegin <= len("montant de") {
		return 0
	}
	indexEnd := strings.Index(line, "€")
	prizeStr := strings.Replace(line[indexBegin:indexEnd], " ", "", -1)
	prizeStr = strings.Replace(prizeStr, "&nbsp;", "", -1)
	prize, err := strconv.Atoi(prizeStr)
	if err != nil {
		// handle error
		l.Error(err)
		os.Exit(2)
	}
	return prize
}

func extractNextLotteryDate(line string) string {
	line = line[len("Le montant du jackpot du prochain tirage du "):]
	index := strings.Index(line, "est")
	return line[:index]
}

// Could be done in a better way
func extractNextLotteryPrize(line string) int {
	indexBegin := strings.Index(line, "est de") + len("est de")
	indexEnd := strings.Index(line, "€")
	prizeStr := strings.Replace(line[indexBegin:indexEnd], " ", "", -1)
	prize, err := strconv.Atoi(prizeStr)
	if err != nil {
		// handle error
		l.Error(err)
		os.Exit(2)
	}
	return prize
}
