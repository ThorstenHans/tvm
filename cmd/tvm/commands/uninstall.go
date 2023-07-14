package commands

import (
	"github.com/spf13/cobra"
)

var all bool = false

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a desired Terraform version",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	uninstallCmd.Flags().BoolVarP(&all, "all", "", false, "Uninstall all Terraform versions")
	rootCmd.AddCommand(uninstallCmd)
}
