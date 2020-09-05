package main

import (
	"os"

	"github.com/halkn/godm"
)

func main() {
	os.Exit(godm.NewCLI(os.Stdout, os.Stderr).Run(os.Args[:]))
}
