package main

import (
	"github.com/K0ng2/bilisubdl/cmd"
)

var version string

func main() {
	cmd.RootCmd.Version = version
	cmd.RootCmd.Execute()
}
