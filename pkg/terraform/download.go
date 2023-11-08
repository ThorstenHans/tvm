package terraform

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ThorstenHans/tvm/pkg/dirs"
	"github.com/briandowns/spinner"
	"golang.org/x/sync/errgroup"

	version "github.com/hashicorp/go-version"
	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/hc-install/src"
)

func Download(ctx context.Context, desiredVersion string) error {
	group, ctx := errgroup.WithContext(ctx)
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Downloading Terraform %s", desiredVersion)
	s.Start()
	group.Go(func() error {
		_, err := doDownload(ctx, desiredVersion)
		return err
	})

	err := group.Wait()
	s.Stop()
	return err
}

func doDownload(ctx context.Context, v string) (string, error) {
	i := install.NewInstaller()
	installDir, err := dirs.GetInstallDir(v)
	if err != nil {
		return "", err
	}
	if err = os.MkdirAll(installDir, 0777); err != nil {
		return "", err
	}

	installable := []src.Installable{
		&releases.ExactVersion{
			Product:                  product.Terraform,
			Version:                  version.Must(version.NewVersion(v)),
			InstallDir:               installDir,
			SkipChecksumVerification: false,
		},
	}

	p, err := i.Install(ctx, installable)

	if err != nil {
		return "", err
	}
	return p, nil
}
