package godm

import (
	"bytes"
	"testing"
)

func TestVersionCommand_exec(t *testing.T) {
	out := new(bytes.Buffer)
	err := new(bytes.Buffer)
	vc := &versionCommand{&command{outWriter: out, errWriter: err}}

	t.Run("exec_normal", func(t *testing.T) {
		if err := vc.exec([]string{"version"}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
