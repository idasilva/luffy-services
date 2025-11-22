package cmd

import (
	"github.com/idasilva/luffy-services/app/modules"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "serve the api",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		
		logger.Info("Executing server command")
		err := modules.Application().Server()
		if err != nil {
			logger.Error("Server execution failed", zap.Error(err))
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
