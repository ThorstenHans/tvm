package commands

import (
	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/spf13/cobra"
)

var listRemoteLimit int = 0

var listRemoteCmd = &cobra.Command{
	Use:     "list-remote",
	Short:   "List all available versions",
	Aliases: []string{"lsr"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return terraform.ListRemote(listRemoteLimit)
	},
}

func init() {
	listRemoteCmd.Flags().IntVarP(&listRemoteLimit, "top", "t", 0, "Show only top n results")
	rootCmd.AddCommand(listRemoteCmd)
}
