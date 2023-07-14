package versionmanager

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Version = string

const (
	xdgConfigDirVar = "XDG_CONFIG_HOME"
	xdgDataDirVar   = "XDG_DATA_HOME"
	tvmFolderName   = "tvm"
	executable      = "terraform"
)

func GetCurrentVersion() (Version, error) {
	return readVersion()
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
	cfgFolder, ok := os.LookupEnv(xdgConfigDirVar)
	var err error
	if !ok {
		cfgFolder, err = os.UserHomeDir()
	}
	if err != nil {
		return "", err
	}
	if err == nil && !ok {
		return path.Join(cfgFolder, fmt.Sprintf(".%s", tvmFolderName), ".tvm"), nil
	}
	return path.Join(cfgFolder, tvmFolderName, ".tvm"), nil
}

func readVersion() (Version, error) {
	cfg, err := getConfigFilePath()
	if err != nil {
		return "", err
	}
	_, err = os.Stat(cfg)
	if err != nil {
		return "Sorry, you haven't installed any Terraform version.", nil
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
	f, ok := os.LookupEnv(xdgDataDirVar)
	var err error
	if !ok {
		f, err = os.UserHomeDir()
	}
	if err != nil {
		return "", nil
	}
	if err == nil && !ok {
		return path.Join(f, fmt.Sprintf(".%s", tvmFolderName), executable), nil
	}
	return path.Join(f, tvmFolderName, executable), nil
}

func getVersionBinaryPath(v Version) (string, error) {
	f, ok := os.LookupEnv(xdgDataDirVar)
	var err error
	if !ok {
		f, err = os.UserHomeDir()
	}
	if err != nil {
		return "", nil
	}
	vDir := strings.ReplaceAll(v, ".", "_")
	if err == nil && !ok {
		return path.Join(f, fmt.Sprintf(".%s", tvmFolderName), vDir, executable), nil
	}
	return path.Join(f, tvmFolderName, vDir, executable), nil
}

func removeLink(p string) error {
	if _, err := os.Stat(p); err == nil {
		return os.Remove(p)
	}
	return nil
}
