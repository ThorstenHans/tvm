package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "tvm",
	Short: "Terraform Version Manager",
}

func Execute() error {
	return rootCmd.Execute()
}
