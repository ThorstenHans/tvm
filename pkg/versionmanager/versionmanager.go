package versionmanager

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/ThorstenHans/tvm/pkg/dirs"
)

type Version = string

const (
	executable = "terraform"
	product    = "Terraform"
	rcFile     = ".terraform-version"
)

func GetCurrentVersion() (Version, error) {
	return readVersion()
}

func PinVersion(v Version) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(dir, rcFile), []byte(fmt.Sprintf("%s\n", v)), 0666)
}
func PinCurrentVersion() error {
	v, err := GetCurrentVersion()
	if err != nil {
		return fmt.Errorf("could not determine current %s version", product)
	}
	return PinVersion(v)

}

func GetPinnedVersion() (Version, bool) {
	dir, err := os.Getwd()
	if err != nil {
		return "", false
	}
	v, err := os.ReadFile(path.Join(dir, rcFile))
	if err != nil {
		return "", false
	}
	ver := string(v)
	ver = strings.TrimSpace(ver)
	return ver, true
}

func UseVersion(v Version) error {
	err := linkBinary(v)
	if err != nil {
		return err
	}
	return storeVersion(v)
}

func storeVersion(v Version) error {
	cfg, err := getConfigFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(cfg, []byte(v), 0666)
}

func getConfigFilePath() (string, error) {
	dir, err := dirs.GetTerraformVersionManagerFolder()
	if err != nil {
		return "", err
	}
	return path.Join(dir, ".tvm"), nil
}

func readVersion() (Version, error) {
	cfg, err := getConfigFilePath()
	if err != nil {
		return "", err
	}
	_, err = os.Stat(cfg)
	if err != nil {
		return "No versions of Terraform installed yet.", nil
	}
	d, err := os.ReadFile(cfg)
	if err != nil {
		return "", err
	}
	return string(d), nil

}

func linkBinary(v Version) error {
	target, err := getExecutableFilePath()
	if err != nil {
		return err
	}
	versionBinaryPath, err := getVersionBinaryPath(v)
	if err != nil {
		return err
	}
	err = removeLink(target)
	if err != nil {
		return err
	}
	return os.Link(versionBinaryPath, target)
}

func getExecutableFilePath() (string, error) {
	f, err := dirs.GetTerraformVersionManagerFolder()
	if err != nil {
		return "", nil
	}
	return path.Join(f, executable), nil
}

func getVersionBinaryPath(v Version) (string, error) {
	f, err := dirs.GetInstallDir(v)
	if err != nil {
		return "", nil
	}
	return path.Join(f, executable), nil
}

func removeLink(p string) error {
	if _, err := os.Stat(p); err == nil {
		return os.Remove(p)
	}
	return nil
}
