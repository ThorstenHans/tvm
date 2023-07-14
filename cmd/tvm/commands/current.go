package commands

import (
	"fmt"
	"github.com/ThorstenHans/tvm/pkg/versionmanager"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:     "current",
	Short:   "Print the currently active Terraform version",
	Aliases: []string{"cur"},
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := versionmanager.GetCurrentVersion()
		if err != nil {
			return err
		}
		fmt.Println(v)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
