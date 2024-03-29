package terraform

import (
	"os"
	"strings"

	"github.com/ThorstenHans/tvm/pkg/dirs"
)

func ListInstalled(limit int) error {
	tvmFolder, err := dirs.GetTerraformVersionManagerFolder()
	if err != nil {
		return nil
	}

	entries, err := os.ReadDir(tvmFolder)
	if err != nil {
		return nil
	}
	versions := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			versions = append(versions, strings.ReplaceAll(entry.Name(), "_", "."))
		}
	}
	sortVersions(versions)
	printVersions(versions, limit)
	return nil
}
