package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
			fmt.Println("info:", err.Error())
			return
		}
		fmt.Println("config created at:", path)
	},
}

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View Codewise config",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadConfig()
		if err != nil {
			fmt.Println("info:", err.Error())
			return
		}

		out, err := yaml.Marshal(cfg)
		if err != nil {
			fmt.Println("info: failed to render config")
			return
		}

		fmt.Println(string(out))
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configViewCmd)
	rootCmd.AddCommand(configCmd)
}
