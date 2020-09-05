package godm

import "fmt"

const version = "0.1.0"

type versionCommand struct {
	*command
}

func (vc *versionCommand) exec(args []string) error {
	vc.out(fmt.Sprintf("godm's version: %s", version))
	return nil
}
