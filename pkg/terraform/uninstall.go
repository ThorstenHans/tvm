package terraform

import (
	"os"

	"github.com/ThorstenHans/tvm/pkg/dirs"
)

func Uninstall(v string) error {
	d, err := dirs.GetInstallDir(v)
	if err != nil {
		return err
	}
	if err = os.RemoveAll(d); err != nil {
		return err
	}
	return nil
}

func IsVersionInstalled(v string) (bool, error) {
	d, err := dirs.GetInstallDir(v)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(d)
	if err != nil {
		return false, nil
	}
	return true, nil

}
