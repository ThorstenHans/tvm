package commands

import (
	"fmt"

	"github.com/ThorstenHans/tvm/pkg/printer"
	"github.com/spf13/cobra"
)

var p printer.Printer

var rootCmd = &cobra.Command{
	Use:   "tvm",
	Short: "Terraform Version Manager",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		p = printer.NewPrinter(verbose)
	},
}

func SetVersionMetadata(version string, commit string, date string) {
	rootCmd.Version = version
	sha := ""
	if len(commit) >= 7 {
		sha = fmt.Sprintf(" (%s)", commit[:7])
	}
	rootCmd.SetVersionTemplate(fmt.Sprintf("tvm version {{.Version}}%s\n", sha))
}

func init() {
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose output")
}
func Execute() error {
	return rootCmd.Execute()
}
