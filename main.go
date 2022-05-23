package main

import (
	"log"

	"github.com/K0ng2/bilisubdl/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
