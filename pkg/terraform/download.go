package terraform

import (
	"context"
	"fmt"
	version "github.com/hashicorp/go-version"
	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/hc-install/src"
	"os"
	"path"
	"strings"
)

const (
	xdgDataDirVar = "XDG_DATA_HOME"
	tvmFolderName = "tvm"
)

func getTerraformVersionManagerFolder() (string, error) {
	p, ok := os.LookupEnv("xdgDataDirVar")
	var err error
	if !ok {
		p, err = os.UserHomeDir()
	}
	if err != nil {
		return "", nil
	}
	if err == nil && !ok {
		return path.Join(p, fmt.Sprintf(".%s", tvmFolderName)), nil
	}
	return path.Join(p, tvmFolderName), nil
}

func getInstallDir(version string) (string, error) {
	tvm, err := getTerraformVersionManagerFolder()
	if err != nil {
		return "", err
	}
	return path.Join(tvm, strings.ReplaceAll(version, ".", "_")), nil
}
func Download(ctx context.Context, desiredVersion string) (string, error) {
	i := install.NewInstaller()
	installDir, err := getInstallDir(desiredVersion)
	if err != nil {
		return "", err
	}
	if err = os.MkdirAll(installDir, 0777); err != nil {
		return "", err
	}

	installable := []src.Installable{
		&releases.ExactVersion{
			Product:    product.Terraform,
			Version:    version.Must(version.NewVersion(desiredVersion)),
			InstallDir: installDir,
		},
	}

	p, err := i.Install(ctx, installable)
	if err != nil {
		return "", err
	}
	return p, nil
}
