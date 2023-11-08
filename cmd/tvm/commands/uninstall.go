package commands

import (
	"os"

	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/ThorstenHans/tvm/pkg/versionmanager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var all bool = false

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Uninstall a desired Terraform version",
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"del", "rm"},
	Run: func(_cmd *cobra.Command, args []string) {
		desired := args[0]
		i, err := terraform.IsVersionInstalled(desired)
		if err != nil {
			color.Red("Can't determine if version is installed: %s", desired)
			os.Exit(1)
		}

		if !i {
			color.Red("Version %s not installed", desired)
			os.Exit(1)
		}
		cur, err := versionmanager.GetCurrentVersion()
		if err != nil {
			color.Red("Could not determine current version")
			os.Exit(1)
		}
		if desired == cur {
			color.Red("Cannot uninstall currently active version")
			os.Exit(1)
		}
		err = terraform.Uninstall(desired)
		if err != nil {
			color.Red("Error while uninstalling version %s\n%s", desired, err)
			os.Exit(1)
		}
		color.Green("Successfully uninstalled version %s", desired)
		os.Exit(0)
	},
}

func init() {
	uninstallCmd.Flags().BoolVarP(&all, "all", "", false, "Uninstall all Terraform versions")
	rootCmd.AddCommand(uninstallCmd)
}
