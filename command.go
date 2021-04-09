package godm

import (
	"fmt"
	"io"
	"log"

	"github.com/halkn/godm/installer"
)

type execer interface {
	exec([]string) error
}

type command struct {
	outWriter io.Writer
	errWriter io.Writer
	config
}

func (c *command) out(v ...interface{}) {
	fmt.Fprintln(c.outWriter, v...)
}

func (c *command) log(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func newSubCommand(sub string, cmd *command) execer {
	switch sub {
	case "deps":
		return &depsCommand{command: cmd}
	case "edit":
		return &editCommand{command: cmd}
	case "install":
		return &installCommand{command: cmd, Installer: installer.NewCommandInstaller()}
	case "update":
		log.Println("update command is not yet implemented")
		return nil
	case "clean":
		log.Println("clean command is not yet implemented")
		return nil
	case "version":
		return &versionCommand{command: cmd}
	default:
		return nil
	}
}
