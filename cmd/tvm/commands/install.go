package commands

import (
	"context"
	"fmt"
	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a desired Terraform version",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		v := args[0]
		p, err := terraform.Download(ctx, v)
		fmt.Println(p)
		return err
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
