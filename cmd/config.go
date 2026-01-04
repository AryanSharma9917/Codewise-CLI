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
		fmt.Println("✅ Config created at:", path)
	},
}

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View Codewise config",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := config.ReadConfig()
		if err != nil {
			fmt.Println("ℹ️", err.Error())
			return
		}
		fmt.Println(string(data))
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configViewCmd)
	rootCmd.AddCommand(configCmd)
}
