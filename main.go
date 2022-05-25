package main

import (
	"log"

	"github.com/K0ng2/bilisubdl/cmd"
)

var version string

func main() {
	cmd.RootCmd.Version = version
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
