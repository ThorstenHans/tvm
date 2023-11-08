package commands

import (
	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List installed Terraform versions (limited to 40 most recent releases by default)",
	Aliases: []string{"ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		limit := 40
		all, _ := cmd.Flags().GetBool("all")
		if all {
			limit = 0
		}
		return terraform.ListInstalled(limit)
	},
}

func init() {
	listCmd.Flags().Bool("all", false, "Show all releases")
	rootCmd.AddCommand(listCmd)
}
