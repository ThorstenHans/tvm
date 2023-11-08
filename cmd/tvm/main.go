package main

import (
	"fmt"
	"os"

	"github.com/ThorstenHans/tvm/cmd/tvm/commands"
)

var (
	version = "some"
	commit  = "some"
	date    = "some"
)

func init() {
	commands.SetVersionMetadata(version, commit, date)
}
func main() {

	if err := commands.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
