package godm

import (
	"os"
	"os/exec"
)

type editCommand struct {
	*command
}

func (ec *editCommand) exec(args []string) error {

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command("sh", "-c", editor, "-c '"+os.ExpandEnv(ec.Installpath)+"'")
	cmd.Stdin = os.Stdin
	cmd.Stdout = ec.outWriter

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}
