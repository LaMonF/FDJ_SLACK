package main

import (
	"github.com/LaMonF/FDJ_SLACK/service"
)

func main() {
	startServer()
}

func startServer() {
	service.SetUpCron()
	service.SetUpServer()
}

