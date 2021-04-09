// Package godm ...
package godm

import (
	"fmt"
	"io"
	"log"
)

const (
	cmdPattern = "deps|install|update|clean|version"
)

type CLI struct {
	outStream, errStream io.Writer
}

func NewCLI(out, err io.Writer) *CLI {
	return &CLI{outStream: out, errStream: err}
}

func (c *CLI) Run(args []string) int {
	log.SetOutput(c.errStream)
	log.SetPrefix("[godm] ")
	log.SetFlags(log.Ltime)

	if len(args) < 2 {
		fmt.Fprintf(c.errStream, "expected subcommand [%s]\n", cmdPattern)
		return 1
	}

	cfg, err := loadConfig()
	if err != nil {
		log.Printf("failed to load config: %v", err)
		return 1
	}

	subcmd := newSubCommand(args[1], &command{outWriter: c.outStream, errWriter: c.errStream, config: *cfg})
	if subcmd == nil {
		fmt.Fprintf(c.errStream, "expected subcommand [%s]\n", cmdPattern)
		return 1
	}

	if err := subcmd.exec(args[1:]); err != nil {
		log.Printf("failed to %s exec\n%v", args[1], err)
		return 1
	}

	return 0
}
