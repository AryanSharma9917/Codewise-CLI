package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Codewise configuration",
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Codewise config file",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := config.InitConfig()
		if err != nil {
			fmt.Println("ℹ️", err.Error())
			return
		}
		fmt.Println("Config created at:", path)
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	rootCmd.AddCommand(configCmd)
}
