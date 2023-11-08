package commands

import (
	"os"

	"github.com/ThorstenHans/tvm/pkg/terraform"
	"github.com/ThorstenHans/tvm/pkg/versionmanager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var explicitVersion string

var pinCmd = &cobra.Command{
	Use:     "pin",
	Short:   "Track the desired Terraform version in a .terraform-version file",
	Long:    "Use `tvm pin` to track the desired Terraform version. The command will \neither take the currently active Terraform version,or the version specified \nusing the `--explicit (-e)` flag. The version is stored in a .terraform-version file.\nThe command will create the .terraform-version file if non-existent.",
	Example: " - `tvm pin` pins the currently active version of Terraform\n - `tvm pin -e 1.0.0` pins version 1.0.0 of Terraform\n",
	Run: func(_cmd *cobra.Command, _args []string) {

		var v string
		if len(explicitVersion) > 0 {
			v = explicitVersion
			// check if v is installed
			ok, err := terraform.IsVersionInstalled(v)
			if err != nil {
				color.Red("Could not determine if version %s is installed", v)
				os.Exit(1)
			}
			if !ok {
				color.Red("Version %s is not installed", v)
				color.Yellow("Install it using `tvm install %s`", v)
				os.Exit(1)
			}
		} else {
			c, err := versionmanager.GetCurrentVersion()
			if err != nil {
				color.Red("Could not determine current version")
				os.Exit(1)
			}
			v = c
		}
		err := versionmanager.PinVersion(v)
		if err != nil {
			color.Red("Error while pinning Terraform version %s\n", v)
			os.Exit(1)
		}
		color.Green("Successfully pinned version %s in .terraform-version file", v)
	},
}

func init() {
	pinCmd.Flags().StringVarP(&explicitVersion, "explicit", "e", "", "Pin to a explicit version of Terraform")
	rootCmd.AddCommand(pinCmd)
}
