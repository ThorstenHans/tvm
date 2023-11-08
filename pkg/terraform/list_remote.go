package terraform

import (
	"encoding/json"
	"net/http"
)

const url = "https://releases.hashicorp.com/terraform/index.json"

type Releases struct {
	Name     string             `json:"name"`
	Versions map[string]Version `json:"versions"`
}

type Version struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func ListRemote(limit int) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	var releases Releases
	err = json.NewDecoder(r.Body).Decode(&releases)
	if err != nil {
		return err
	}
	versions := make([]string, 0)
	for _, v := range releases.Versions {
		versions = append(versions, v.Version)
	}
	sortVersions(versions)
	printVersions(versions, limit)
	return nil
}
