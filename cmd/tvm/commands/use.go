package commands

import (
	"fmt"
	"github.com/ThorstenHans/tvm/pkg/versionmanager"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch to a desired Terraform version",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		v := args[0]
		fmt.Println(v)
		return versionmanager.UseVersion(v)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
