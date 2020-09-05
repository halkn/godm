package godm

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func cleanupDir(dest string) {
	if err := os.RemoveAll(dest); err != nil {
		fmt.Println(err)
	}
}

func isNotExitDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// not exit
		return true
	}
	// exit
	return false
}

func TestDepsCommand_exec(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)

	pwd, _ := os.Getwd()
	testdir := filepath.Join(pwd, "testdata", "depsCommandTest")
	reset := setTestEnv("GODM_TEST_DEPS_COMMAND_TEST", testdir)
	defer reset()
	cleanupDir(testdir)

	pattern := map[string]struct {
		cfg      *config
		args     []string
		out      string
		err      string
		makeddir []string
		isErr    bool
	}{
		"parse error": {
			cfg: &config{
				Dirs: []string{"test/dir", "test/dir2"},
			},
			args:     []string{"deps", "--failllllllllllllllll"},
			out:      "",
			err:      "flag provided but not defined",
			makeddir: nil,
			isErr:    true,
		},
		"list": {
			cfg: &config{
				Dirs: []string{"test/dir", "test/dir2"},
			},
			args:     []string{"deps", "--list"},
			out:      "test/dir\ntest/dir2\n",
			err:      "",
			makeddir: nil,
			isErr:    false,
		},
		"ok": {
			cfg: &config{
				Dirs: []string{
					filepath.Join(testdir, "1"),
					"${GODM_TEST_DEPS_COMMAND_TEST}/2",
					"$GODM_TEST_DEPS_COMMAND_TEST/3",
					"$GODM_TEST_DEPS_COMMAND_TEST/4/5",
				},
			},
			args: []string{"deps", ""},
			out:  "",
			err:  "",
			makeddir: []string{
				filepath.Join(testdir, "1"),
				filepath.Join(testdir, "2"),
				filepath.Join(testdir, "3"),
				filepath.Join(testdir, "4", "5"),
			},
			isErr: false,
		},
	}

	for k, tt := range pattern {
		t.Run(k, func(t *testing.T) {
			cmd := &depsCommand{
				command: &command{
					outWriter: outStream,
					errWriter: errStream,
					config:    *tt.cfg,
				},
			}
			err := cmd.exec(tt.args)
			if tt.isErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if !strings.Contains(err.Error(), tt.err) {
					t.Fatalf("got: %v, want: %v", err.Error(), tt.err)
				}
			}
			if tt.out != "" {
				if outStream.String() != tt.out {
					t.Fatalf("got: %v, want: %v", outStream.String(), tt.out)
				}
			}
			if tt.makeddir != nil {
				for _, v := range tt.makeddir {
					if isNotExitDir(v) {
						t.Fatalf("failed to makedir: %s", v)
					}
				}
			}
			outStream.Reset()
			errStream.Reset()
		})
	}

	cleanupDir(testdir)
}
