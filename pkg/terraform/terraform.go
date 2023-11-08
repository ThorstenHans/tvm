package terraform

import (
	"fmt"
	"sort"
)

func sortVersions(v []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(v)))
}
func limitVersions(v []string, limit int) []string {
	if limit > 0 && len(v) > limit {
		return v[:limit]
	}
	return v
}
func printVersions(v []string, limit int) {
	limited := limitVersions(v, limit)
	for _, ver := range limited {
		fmt.Println(ver)
	}
}
