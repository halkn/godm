package godm

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	pwd, _ := os.Getwd()

	reset := setTestEnv("HOME", pwd)
	defer reset()

	pattern := map[string]struct {
		confdir string
		want    *config
		isErr   bool
	}{
		"xdg_ok": {
			confdir: filepath.Join(pwd, "testdata"),
			want: &config{
				Installpath: "$HOME/.dotfiles",
				Dirs:        []string{"path/to/test", "path/to/test1"},
			},
			isErr: false,
		},
		"xdg_ng": {
			confdir: filepath.Join(pwd, "testdataF"),
			want:    nil,
			isErr:   true,
		},
		"home": {
			confdir: "",
			want:    nil,
			isErr:   true,
		},
	}

	for k, tt := range pattern {
		t.Run(k, func(t *testing.T) {
			reset := setTestEnv("XDG_CONFIG_HOME", tt.confdir)
			defer reset()
			got, err := loadConfig()
			if tt.isErr {
				if err == nil {
					t.Fatal("expected error")
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
