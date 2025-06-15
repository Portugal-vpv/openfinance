package main

import (
	"openfinance/cmd"
	"openfinance/configuration/logger"
)

func main() {
	logger.NewLogger()
	cmd.StartHttpServer()
}
