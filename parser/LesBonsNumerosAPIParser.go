package parser

import (
	l "FDJ_SLACK/log"
	"FDJ_SLACK/model"
	"FDJ_SLACK/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/etree"
)

type LesBonsNumerosAPIParser struct {
}

func NewParser() *LesBonsNumerosAPIParser {
	s := &LesBonsNumerosAPIParser{}
	return s
}

func (p *LesBonsNumerosAPIParser) GetAndParseData(apiURL string) []model.LotteryResult {
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

	l.Info("Parsing data")

	//Create new XML Document to go through results
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(contents); err != nil {
		panic(err)
	}

	var result []model.LotteryResult

	//For each item (using XPath to get description)
	for _, t := range doc.FindElements("//item/description") {
		l.Info("Found item")
		for _, child := range t.Child {
			if c, ok := child.(*etree.CharData); ok {

				// Manage only not empty strings
				if !isWhitespace(c.Data) {

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
					if len(resultByLine) > 4 { // When results are not up to date len(result_line_list) == 4
						numberWinner = extractNumberWinner(resultByLine[11])
						winnerPrize = extractWinnerPrize(resultByLine[11])
						nextLotteryDate = extractNextLotteryDate(resultByLine[12])
						nextLotteryPrize = extractNextLotteryPrize(resultByLine[12])
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

// TODO: Extract Winner prize
func extractWinnerPrize(line string) int {
	// prize = re.findall(r'\d+', line.replace(" ",""))
	// return 0 if len(prize) == 0 else int(prize[0])
	l.Warning("TODO ", line)
	return 0
}

func extractNextLotteryDate(line string) string {
	line = line[len("Le montant du jackpot du prochain tirage du "):]
	index := strings.Index(line, "est")
	return line[:index]
}

// TODO: Extract Next Lottery prize
func extractNextLotteryPrize(line string) int {
	// Detected numbers in the line and filter when greater than 6 characters (millions)
	// return int(list(prize for prize in re.findall(r'\d+', line.replace(" ", "")) if len(prize) > 6)[0]) // too complex
	l.Warning("TODO ", line)
	return 0
}

// isWhitespace returns true if the byte slice contains only
// whitespace characters.
func isWhitespace(s string) bool {
	for i := 0; i < len(s); i++ {
		if c := s[i]; c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			return false
		}
	}
	return true
}
