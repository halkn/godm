package godm

import (
	"flag"
	"os"
)

type depsCommand struct {
	*command
}

func (dc *depsCommand) exec(args []string) error {

	var showList bool

	flg := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flg.SetOutput(dc.errWriter)

	flg.BoolVar(&showList, "list", false, "Print setup directory list.")
	if err := flg.Parse(args[1:]); err != nil {
		return err
	}

	if showList {
		for _, v := range dc.config.Dirs {
			dc.out(os.ExpandEnv(v))
		}
		return nil
	}

	for _, v := range dc.config.Dirs {
		if err := os.MkdirAll(os.ExpandEnv(v), 0755); err != nil {
			dc.log("failed to create dir[%s]: %v", os.ExpandEnv(v), err)
		}
	}

	return nil
}
