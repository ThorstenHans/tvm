package commands

import (
	"fmt"
	"os"

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
		var err error
		// todo refactor this to load version and store it locally this would allow presenting more meaningful messages
		if len(explicitVersion) > 0 {
			//todo check if version is installed
			err = versionmanager.PinVersion(explicitVersion)
		} else {
			err = versionmanager.PinCurrentVersion()
		}
		if err != nil {
			color.Red("Sorry, could not pin Terraform version")
			fmt.Println(err)
			os.Exit(1)
		}
		color.Green("Successfully pinned Terraform version in .terraform-version file")
	},
}

func init() {
	pinCmd.Flags().StringVarP(&explicitVersion, "explicit", "e", "", "Pin to a explicit version of Terraform")
	rootCmd.AddCommand(pinCmd)
}
