package cmd

import (
	"log"
	"os"

	"github.com/aryansharma9917/Codewise-CLI/cmd/docker"
	"github.com/aryansharma9917/Codewise-CLI/cmd/kubernetes"
	"github.com/spf13/cobra"
)

// rootCmd is the root command for Codewise-CLI
var rootCmd = &cobra.Command{
	Use:   "Codewise-CLI [command]",
	Short: "Do all your tedious tasks with a single command",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		checkNilErr(err)
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Subcommands for the root command
	rootCmd.AddCommand(encodeCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(yamlToJsonCmd)
	rootCmd.AddCommand(jsonToYamlCmd)
	rootCmd.AddCommand(docker.DockerCmd)
	rootCmd.AddCommand(keyValueToJsonCmd)
	rootCmd.AddCommand(kubernetes.KubernetesCmd)

}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
