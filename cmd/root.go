package cmd

import (
	"core/infra/config"
	"core/infra/conn"
	"core/infra/logger"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "Golang Boilerplate",
		Short: "Boilerplate",
		Long:  "To use different microservice as a starter template",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(seedCmd)
}

// Execute executes the root command
func Execute() {
	config.LoadConfig()
	conn.ConnectDb()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.Info("about to start the application")
}
