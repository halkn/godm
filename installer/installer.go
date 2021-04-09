// Package installer is interface to install tools
package installer

import (
	"fmt"
	"os/exec"
)

type Installer interface {
	InstallTool(command []string) error
}

type CommandInstaller struct{}

func NewCommandInstaller() Installer {
	return &CommandInstaller{}
}

func (ti *CommandInstaller) InstallTool(command []string) error {
	if _, err := exec.LookPath(command[0]); err != nil {
		return err
	}

	if out, err := exec.Command(command[0], command[1:]...).CombinedOutput(); err != nil {
		return fmt.Errorf("%w, message: %s", err, string(out))
	}

	return nil
}
