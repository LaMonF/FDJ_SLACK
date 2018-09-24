import re
import logging
import xml.etree.ElementTree as ET

from urllib.request import Request, urlopen

from Utils import Utils, NumbersInFrench
from Result import Result

logger = logging.getLogger("FDJ_SLACK")


class LesBonsNumerosAPIParser(object):

    API_URL = "https://www.lesbonsnumeros.com/loto/rss.xml"

    def __init__(self):
        self.string_response = None
        self.xml_node = None

    def get_data(self):

        request = Request(self.API_URL)
        self.string_response = urlopen(request).read().decode()
        # TODO check response code
        self.xml_node = ET.fromstring(self.string_response)
        logger.debug("[LesBonsNumerosAPIParser] get_data: REPONSE >> " + self.string_response)

    def parse_data(self):
        list_results = list()
        for item_idx, item in enumerate(self.xml_node[0].findall('item')):
            result_text_data = item[3].text
            result_text_data = Utils.cleanhtml(result_text_data)
            result_line_list = result_text_data.split("\n")

            date = LesBonsNumerosAPIParser.__extract_date_lottery__(result_line_list[1])
            balls = LesBonsNumerosAPIParser.__extract_balls__(result_line_list[2])
            lucky_ball = LesBonsNumerosAPIParser.__extract_lucky_ball__(result_line_list[3])
            number_winner = LesBonsNumerosAPIParser.__extract_number_winner__(result_line_list[11])
            winner_prize = LesBonsNumerosAPIParser.__extract_winner_prize__(result_line_list[11])
            next_lottery_date = LesBonsNumerosAPIParser.__extract_next_lottery_date__(result_line_list[12])
            next_lottery_prize = LesBonsNumerosAPIParser.__extract_next_lottery_prize__(result_line_list[12])

            result = Result(string_date=date,
                            balls=balls,
                            lucky_ball=lucky_ball,
                            number_winner=number_winner,
                            winner_prize=winner_prize,
                            next_lottery_string_date=next_lottery_date,
                            next_lottery_prize=next_lottery_prize)

            list_results.append(result)
        logger.debug("[LesBonsNumerosAPIParser] parse_data: LIST_RESUlTS >> " + str(list_results))
        return list_results

    @staticmethod
    def __extract_date_lottery__(line):
        return line[len("Résultats du "):]

    @staticmethod
    def __extract_balls__(line):
        line = line[len("Numéros : "):]
        return line.split(" - ")

    @staticmethod
    def __extract_lucky_ball__(line):
        return line[len("Numéro Chance : ")]

    @staticmethod
    def __extract_number_winner__(line):
        winner_number = -1
        if "Le jackpot n'a pas été remporté lors de ce tirage !" in line:
            winner_number = 0
            line = line.lower()
        elif NumbersInFrench.ONE in line:
            winner_number = 1
        elif NumbersInFrench.TWO in line:
            winner_number = 2
        elif NumbersInFrench.THREE in line:
            winner_number = 3
        elif NumbersInFrench.FOUR in line:
            winner_number = 4
        elif NumbersInFrench.FIVE in line:
            winner_number = 5
        elif NumbersInFrench.SIX in line:
            winner_number = 6
        elif NumbersInFrench.SEVEN in line:
            winner_number = 7
        elif NumbersInFrench.EIGHT in line:
            winner_number = 8
        elif NumbersInFrench.NINE in line:
            winner_number = 9
        elif NumbersInFrench.TEN in line:
            winner_number = 10
        return winner_number

    @staticmethod
    def __extract_winner_prize__(line):
        prize = re.findall(r'\d+', line.replace(" ",""))
        return 0 if len(prize) == 0 else int(prize[0])

    @staticmethod
    def __extract_next_lottery_date__(line):
        line = line[len("Le montant du jackpot du prochain tirage du "):]
        line = " ".join(line.split(" ")[:4])[:-3] # remove the 'est' attached to the year coming from the API.
        return line

    @staticmethod
    def __extract_next_lottery_prize__(line):
        # Detected numbers in the line and filter when greater than 6 characters (millions)
        return int(list(prize for prize in re.findall(r'\d+', line.replace(" ", "")) if len(prize) > 6)[0]) # too complex
