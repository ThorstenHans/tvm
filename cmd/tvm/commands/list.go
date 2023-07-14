package commands

import (
	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/spf13/cobra"
)

var listLimit int = 0
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List installed Terraform versions",
	Aliases: []string{"ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return terraform.ListInstalled(listLimit)
	},
}

func init() {
	listCmd.Flags().IntVarP(&listLimit, "top", "t", 0, "Show only top n results")
	rootCmd.AddCommand(listCmd)
}
