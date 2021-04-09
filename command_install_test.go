package godm

import (
	"bytes"
	"errors"
	"testing"
)

type dummyInstaller struct{}

func (di *dummyInstaller) InstallTool(command []string) error {
	if command[0] == "error" {
		return errors.New("InstallTool error test")
	}
	return nil
}

func TestInstallCommand_exec(t *testing.T) {
	out := new(bytes.Buffer)
	err := new(bytes.Buffer)

	cmd := &command{
		outWriter: out,
		errWriter: err,
		config: config{
			Installpath: "",
			Dirs:        []string{},
			Tools:       [][]string{},
		},
	}

	ic := &installCommand{
		command:   cmd,
		Installer: &dummyInstaller{},
	}

	pattern := map[string]struct {
		tools [][]string
		err   error
	}{
		"ok": {
			tools: [][]string{
				{"go", "install", "path"},
				{"ok", "path2"},
			},
			err: nil,
		},
		"ng": {
			tools: [][]string{
				{"error", "command"},
			},
			err: errors.New("InstallTool error test"),
		},
	}

	for k, tt := range pattern {
		t.Run(k, func(t *testing.T) {
			ic.command.Tools = tt.tools
			goterr := ic.exec([]string{""})
			if tt.err != nil {
				if goterr == nil {
					t.Fatalf("want err")
				}
				if tt.err.Error() != goterr.Error() {
					t.Fatalf("got: %s, want: %s", goterr, tt.err)
				}
			} else {
				if goterr != nil {
					t.Fatalf("Want error is nil but \n%v", goterr)
				}
			}
			out.Reset()
			err.Reset()
		})
	}
}
