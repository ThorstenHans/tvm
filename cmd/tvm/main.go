package main

import (
	"fmt"
	"os"

	"github.com/ThorstenHans/tvm/cmd/tvm/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
