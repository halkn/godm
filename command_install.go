package godm

import (
	"github.com/halkn/godm/installer"
)

type installCommand struct {
	*command
	installer.Installer
}

func (ic *installCommand) exec(args []string) error {
	for _, tool := range ic.Tools {
		if err := ic.InstallTool(tool); err != nil {
			return err
		}
	}
	return nil
}
