package commands

import (
	"context"
	"os"

	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/fatih/color"
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
			color.Red("Error while installing version %s\n%s", v, err)
			os.Exit(1)
		}
		color.Green("Terraform version %s installed successfully.\nYou can switch to it with the `tvm use %s` command\n", v, v)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
