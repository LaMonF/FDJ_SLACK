package utils

import (
	"os"
	"regexp"
)

const ONE = "un"
const TWO = "deux"
const THREE = "trois"
const FOUR = "quatre"
const FIVE = "cinq"
const SIX = "six"
const SEVEN = "sept"
const EIGHT = "huit"
const NINE = "neuf"
const TEN = "dix"
const HUNDRED = "cent"
const THOUSAND = "mille"
const MILLION = "million"
const BILLION = "milliard"

// It defines the winning rank which determines the amount of winning money.
// cf : https://github.com/LaMonF/FDJ_SLACK/issues/10
type WIN_RANK float64
const (
	RANK_1 WIN_RANK = 2000000
	RANK_2 WIN_RANK = 100000
	RANK_3 WIN_RANK = 1000
	RANK_4 WIN_RANK = 500
	RANK_5 WIN_RANK = 50
	RANK_6 WIN_RANK = 20
	RANK_7 WIN_RANK = 10
	RANK_8 WIN_RANK = 5
	RANK_9 WIN_RANK = 2.20
	RANK_0 WIN_RANK = 0
)

func CleanHTML(rawHTML string) string {
	r := "<.*?>"
	re := regexp.MustCompile(r)
	return re.ReplaceAllString(rawHTML, "")
}

// IsWhitespace returns true if the byte slice contains only
// whitespace characters.
func IsWhitespace(s string) bool {
	for i := 0; i < len(s); i++ {
		if c := s[i]; c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			return false
		}
	}
	return true
}

func TestEq(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ArrayNumberSameOccurence(array1 []int, array2 []int) int {
	occurence := 0
	for _, a := range array2 {
		if Contains(array1, a) {
			occurence++
		}
	}
	return occurence
}