package commands

import (
	"os"

	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/spf13/cobra"
)

var listRemoteCmd = &cobra.Command{
	Use:     "list-remote",
	Short:   "List available Terraform versions (limited to 40 most recent releases by default)",
	Aliases: []string{"lsr"},
	Run: func(cmd *cobra.Command, args []string) {
		limit := 40
		all, _ := cmd.Flags().GetBool("all")
		if all {
			limit = 0
		}
		err := terraform.ListRemote(limit)
		if err != nil {
			p.Error(err, "Error while listing remote versions\n")
			os.Exit(1)
		}
	},
}

func init() {
	listRemoteCmd.Flags().Bool("all", false, "Show all available releases")

	rootCmd.AddCommand(listRemoteCmd)
}
