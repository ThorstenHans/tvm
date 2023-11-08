package commands

import (
	"os"

	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List installed Terraform versions (limited to 40 most recent releases by default)",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		limit := 40
		all, _ := cmd.Flags().GetBool("all")
		if all {
			limit = 0
		}
		err := terraform.ListInstalled(limit)
		if err != nil {
			p.Error(err, "Error while listing installed versions\n")
			os.Exit(1)
		}
	},
}

func init() {
	listCmd.Flags().Bool("all", false, "Show all releases")
	rootCmd.AddCommand(listCmd)
}
