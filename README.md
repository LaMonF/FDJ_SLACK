# FDJ_SLACK
Slack Plugin project to play with the `Française Des Jeux` Loto.


### BUILD

You can build the project using the common go build command :
```go
go build FDJ_SLACK.go
```

### RUN 

The script uses the `SLACK_HOOK_URL` environment variable to define the Slack web-hook.  
```bash
export SLACK_HOOK_URL=https://hooks.slack.com/services/YOUR/API/TOKEN
```

### FEATURES

##### API
Current API_VERSION `V1`

| URL                     | DESCRIPTION                 |
|-------------------------|-----------------------------|
| /API_VERSION/lotoresult | Get the last drawn results. |
| /API_VERSION/balls      | Get the current bet balls.  |
| /API_VERSION/balance    | Get the current balance.    |


##### Balance Management
The balance calculation is based on the official documentation (2018) provided by FDJ. 

You can define a balance management file `balance.fdjSlack` next to your project executable.
It handles `float64` with 2 numbers after coma. (eg: 66.66)


##### Scheduled events

- Every day at 22h15, the last drawn is published on the channel.

- Every MONDAY, WEDNESDAY, SATURDAY at 22h the balance is adjusted based on the last dranw results.

