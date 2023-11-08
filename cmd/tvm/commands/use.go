package commands

import (
	"os"

	"github.com/ThorstenHans/tvm/pkg/versionmanager"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch the actual Terraform version",
	Long: `Switch the actual Terraform version. If a .terraform-version file is present, 
    it will switch to the version listed in the file. Otherwise, the desired 
    version must be passed as an argument.`,

	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, ok := versionmanager.GetPinnedVersion()
		if !ok && len(args) == 0 {
			p.Error(nil, "No local .terraform-version file found! You must provide a version as argument")
			os.Exit(1)
		}
		if !ok {
			v = args[0]
		}
		err := versionmanager.UseVersion(v)
		if err != nil {
			p.Errorf(err, "Error while switching to version %s\n", v)
			os.Exit(1)
		}
		p.Successf("Switched to Terraform %s", v)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
