package installer

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestNewCommandInstaller(t *testing.T) {
	got := NewCommandInstaller()
	want := &CommandInstaller{}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func TestCommanInstaller_InstallTool(t *testing.T) {
	ci := &CommandInstaller{}

	pattern := map[string]struct {
		command []string
		wantErr error
	}{
		"executable_ok": {
			command: []string{"echo", "args1"},
			wantErr: nil,
		},
		"executable_ng": {
			command: []string{"go", "xxx"},
			wantErr: errors.New("exit status 2, message: go xxx: unknown command"),
		},
		"not_executable": {
			command: []string{"notfoundcommand", "args1"},
			wantErr: errors.New("executable file not found"),
		},
	}

	for k, tt := range pattern {
		t.Run(k, func(t *testing.T) {
			goterr := ci.InstallTool(tt.command)
			if tt.wantErr != nil {
				if goterr == nil {
					t.Fatal("want error")
				}
				if !strings.Contains(goterr.Error(), tt.wantErr.Error()) {
					t.Fatalf("\ngot : %s,\nwant: %s", goterr, tt.wantErr)
				}
			} else {
				if goterr != nil {
					t.Fatalf("got error: %v", goterr)
				}
			}
		})
	}
}
