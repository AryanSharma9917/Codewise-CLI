package cmd

import (
"fmt"

"github.com/aryansharma9917/codewise-cli/pkg/env"
"github.com/spf13/cobra"
)

var envListCmd = &cobra.Command{
Use:   "list",
Short: "List environments",
RunE: func(cmd *cobra.Command, args []string) error {
envs, err := env.ListEnvs()
if err != nil {
return LogError(err.Error())
}

if len(envs) == 0 {
fmt.Println("no environments found")
return nil
}

for _, e := range envs {
fmt.Printf("%-10s namespace=%s context=%s\n",
e.Name, e.K8s.Namespace, e.K8s.Context)
}
return nil
},
}

func init() {
envCmd.AddCommand(envListCmd)
}
