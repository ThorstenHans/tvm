package commands

import (
	"os"

	"github.com/ThorstenHans/tvm/pkg/versionmanager"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:     "current",
	Short:   "Print the currently active Terraform version",
	Aliases: []string{"cur"},
	Run: func(cmd *cobra.Command, args []string) {
		v, err := versionmanager.GetCurrentVersion()
		if err != nil {
			p.Error(err, "Could not determine current version")
			os.Exit(1)
		}
		p.Successf("Current Terraform version: %s", v)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
