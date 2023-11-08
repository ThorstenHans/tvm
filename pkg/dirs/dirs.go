package dirs

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const appName = "tvm"

func GetTerraformVersionManagerFolder() (string, error) {
	p, err := os.UserConfigDir()
	if err == nil {
		return path.Join(p, appName), nil
	}
	h, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(h, fmt.Sprintf(".%s", appName)), nil
}

func GetInstallDir(version string) (string, error) {
	tvm, err := GetTerraformVersionManagerFolder()
	if err != nil {
		return "", err
	}
	return path.Join(tvm, strings.ReplaceAll(version, ".", "_")), nil
}
