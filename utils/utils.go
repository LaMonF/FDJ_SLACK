package utils

import (
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

func CleanHTML(rawHTML string) string {
	r := "<.*?>"
	re := regexp.MustCompile(r)
	return re.ReplaceAllString(rawHTML, "")
}
