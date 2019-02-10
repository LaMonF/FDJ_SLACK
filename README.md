[![Build Status](https://travis-ci.com/LaMonF/FDJ_SLACK.svg?branch=master)](https://travis-ci.com/LaMonF/FDJ_SLACK)
[![Coverage Status](https://coveralls.io/repos/github/LaMonF/FDJ_SLACK/badge.svg?branch=master)](https://coveralls.io/github/LaMonF/FDJ_SLACK?branch=master)
# FDJ_SLACK
Slack Plugin project to play with the `Française Des Jeux` Loto.


## BUILD

Use the common `go build` command to make binaries :
```go
go build
```

## RUN 

#### ENV
The script uses the `SLACK_HOOK_URL` environment variable to define the Slack web-hook.  
```bash
export SLACK_HOOK_URL=https://hooks.slack.com/services/YOUR/API/TOKEN
```

#### SETTINGS
The application loads the `settings.yml` configuration file located next to the executable.
```yaml
# Global Setting file for FDJ Slack

# BALANCE
balance_file : balance.fdjSlack     # relative or absolute file path of the balance
bet_price : 2.20                    # bet price

# BET
bet:
  balls: [7, 14, 22, 28, 42]
  bonus: 5

# CRON
cron_post_slack: 0 13 22 * * *      # POST to Slack the last draw result every day at 10.22pm
```

## FEATURES

#### API
Current *API_VERSION* : `1`

| URL                              | DESCRIPTION                  | Note                         |
|----------------------------------|------------------------------|------------------------------|
| `/<API_VERSION>/lotoresult`      | Get the last draw results.   |                              |
| `/<API_VERSION>/balls `          | Get the current bet balls.   |                              |
| `/<API_VERSION>/balance `        | Get the current balance.     |                              |
| `/<API_VERSION>/balance  42.42`  | Set the new balance to 42.42 | `float` after coma 2 numbers |
| `/<API_VERSION>/paytable      `  | Get the winnings table       |                              |


#### Balance Management
The balance calculation is based on the official documentation (2018) provided by FDJ. 

You **MUST** define a balance management **file path** in the `settings.yml`. 
This file will be loaded by the application then updated by the results of the draws.
It handles `float64` with 2 numbers after coma. (eg: 66.66)


#### Scheduled events
- Every MONDAY, WEDNESDAY, SATURDAY at 22h the balance is adjusted based on the last draw results.

