package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tvm",
	Short: "Terraform Version Manager",
}

func SetVersionMetadata(version string, commit string, date string) {
	rootCmd.Version = version
	sha := ""
	if len(commit) >= 7 {
		sha = fmt.Sprintf(" (%s)", commit[:7])
	}
	rootCmd.SetVersionTemplate(fmt.Sprintf("tvm version {{.Version}}%s\n", sha))
}

func Execute() error {
	return rootCmd.Execute()
}
