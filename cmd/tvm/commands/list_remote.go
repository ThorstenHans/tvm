package commands

import (
	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/spf13/cobra"
)

var listRemoteCmd = &cobra.Command{
	Use:     "list-remote",
	Short:   "List available Terraform versions (limited to 40 most recent releases by default)",
	Aliases: []string{"lsr"},
	RunE: func(cmd *cobra.Command, args []string) error {

		limit := 40
		all, _ := cmd.Flags().GetBool("all")
		if all {
			limit = 0
		}
		return terraform.ListRemote(limit)
	},
}

func init() {
	listRemoteCmd.Flags().Bool("all", false, "Show all available releases")

	rootCmd.AddCommand(listRemoteCmd)
}
