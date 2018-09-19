import re


class NumbersInFrench(object):
    ONE = "un"
    TWO = "deux"
    THREE = "trois"
    FOUR = "quatre"
    FIVE = "cinq"
    SIX = "six"
    SEVEN = "sept"
    EIGHT = "huit"
    NINE = "neuf"
    TEN = "dix"
    HUNDRED = "cent"
    THOUSAND = "mille"
    MILLION = "million"
    BILLION = "milliard"


class Utils(object):

    @staticmethod
    def cleanhtml(raw_html):
        cleanr = re.compile('<.*?>')
        cleantext = re.sub(cleanr, '', raw_html)
        return cleantext