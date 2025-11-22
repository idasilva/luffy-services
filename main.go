package main

import (
	"github.com/idasilva/luffy-services/cmd"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	
	logger.Info("Starting luffy-services application...")
	cmd.Execute()
}
