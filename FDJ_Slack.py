#!/usr/bin/env python3
import logging
import os

import requests

from LesBonsNumerosAPIParser import LesBonsNumerosAPIParser

localhost = 'http://localhost:8888'
SLACK_URL = os.getenv('SLACK_HOOK_URL', localhost)

WIN = [7, 14, 22, 28, 42]
LUCK_WIN = 5

logger = logging.getLogger("FDJ_SLACK")


class FDJSlack(object):
    def run(self):
        parser = LesBonsNumerosAPIParser()
        parser.get_data()
        list_result = parser.parse_data()
        for index, result in enumerate(list_result):
            if index == 0: # only first result
                logger.info(result)
                FDJSlack.__post_to_slack__(str(result))
                if result.is_winning(WIN, LUCK_WIN):
                    print("BANCO !!")

    @staticmethod
    def __post_to_slack__(text_to_post):
        headers = {'Content-Type': 'application/json'}
        data = '{"text" : "' + text_to_post + '"}'
        requests.post(SLACK_URL,
                      data=data.encode('utf-8'),
                      headers=headers)


if __name__ == '__main__':
    FDJSlack().run()


