package godm

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestEditCommand_exec(t *testing.T) {

	out := new(bytes.Buffer)
	err := new(bytes.Buffer)
	pwd, _ := os.Getwd()
	testdir := filepath.Join(pwd, "testdata")

	cmd := &command{
		outWriter: out,
		errWriter: err,
		config: config{
			Installpath: testdir,
		},
	}
	ec := &editCommand{command: cmd}

	pattern := map[string]struct {
		editor string
		isErr  bool
	}{
		"ok_pwd": {
			editor: "pwd",
			isErr:  false,
		},
		"ng_vi": {
			editor: "",
			isErr:  false,
		},
	}

	for k, tt := range pattern {
		t.Run(k, func(t *testing.T) {
			reset := setTestEnv("EDITOR", tt.editor)
			defer reset()
			got := ec.exec([]string{"edit"})
			if tt.isErr && got == nil {
				t.Fatal("want error")
			}
		})
	}
}
