package godm

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// test helper
func setTestEnv(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}

func TestNewCLI(t *testing.T) {
	out := new(bytes.Buffer)
	err := new(bytes.Buffer)
	want := &CLI{outStream: out, errStream: err}

	t.Run("normal", func(t *testing.T) {
		got := NewCLI(out, err)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("got: %v, want:%v", got, want)
		}
	})
}

func TestCLI_Run(t *testing.T) {

	pwd, _ := os.Getwd()

	out := new(bytes.Buffer)
	err := new(bytes.Buffer)

	cli := NewCLI(out, err)

	pattern := map[string]struct {
		config string
		args   []string
		cmd    execer
		wantc  int
		errmes string
	}{
		"args error": {
			config: "testdata",
			args:   []string{"godm"},
			wantc:  1,
			errmes: "expected subcommand",
		},
		"config error": {
			config: "testdataF",
			args:   []string{"godm", "version"},
			wantc:  1,
			errmes: "failed to load config",
		},
		"unexpected_subcommand": {
			config: "testdata",
			args:   []string{"godm", "unexpected_subcommand"},
			wantc:  1,
			errmes: "expected subcommand",
		},
		"version": {
			config: "testdata",
			args:   []string{"godm", "version"},
			wantc:  0,
			errmes: "",
		},
		"deps_err": {
			config: "testdata",
			args:   []string{"godm", "deps", "--xxxxxxxxx"},
			wantc:  1,
			errmes: "failed to exec command",
		},
	}

	for k, tt := range pattern {
		t.Run(k, func(t *testing.T) {
			reset := setTestEnv("XDG_CONFIG_HOME", filepath.Join(pwd, tt.config))
			defer reset()
			got := cli.Run(tt.args)
			if got != tt.wantc {
				t.Fatalf("got: %d, want :%d", got, tt.wantc)
			}
			if tt.wantc == 1 {
				if !strings.Contains(err.String(), tt.errmes) {
					t.Fatalf("got: %v, want: %v", err.String(), tt.errmes)
				}
			}
			out.Reset()
			err.Reset()
		})
	}
}
