package commands

import (
	"context"
	"os"

	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a desired Terraform version",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		v := args[0]
		err := terraform.Download(ctx, v)
		if err != nil {
			p.Errorf(err, "Error while installing Terraform version %s\n", v)
			os.Exit(1)
		}
		p.Successf("Terraform version %s installed successfully.", v)
		p.Hintf("You can switch to it with the `tvm use %s` command\n", v)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
