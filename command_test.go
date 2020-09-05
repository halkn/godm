package godm

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestCommand_out(t *testing.T) {

	out := new(bytes.Buffer)
	err := new(bytes.Buffer)
	wantmes := "want message"

	c := &command{outWriter: out, errWriter: err}

	t.Run("normal", func(t *testing.T) {
		c.out(wantmes)
		if out.String() != fmt.Sprintln(wantmes) {
			t.Fatalf("got: %v, want: %v", out.String(), fmt.Sprintln(wantmes))
		}
		out.Reset()
	})

}

func TestCommand_log(t *testing.T) {

	out := new(bytes.Buffer)
	err := new(bytes.Buffer)
	wantlog := "want log"

	c := &command{outWriter: out, errWriter: err}

	log.SetOutput(err)

	t.Run("nomal", func(t *testing.T) {
		c.log(wantlog)
		if !strings.Contains(err.String(), wantlog) {
			t.Fatalf("got: %v, want: %v", err.String(), wantlog)
		}
		err.Reset()
	})
}

func TestNewSubCommand(t *testing.T) {
	out := new(bytes.Buffer)
	err := new(bytes.Buffer)

	log.SetOutput(err)
	c := &command{outWriter: out, errWriter: err}

	pattern := map[string]struct {
		sub  string
		want execer
	}{
		"deps": {
			sub:  "deps",
			want: &depsCommand{command: c},
		},
		"install": {
			sub:  "install",
			want: nil,
		},
		"update": {
			sub:  "update",
			want: nil,
		},
		"clean": {
			sub:  "clean",
			want: nil,
		},
		"version": {
			sub:  "version",
			want: &versionCommand{command: c},
		},
		"nil_unexpected": {
			sub:  "xxxxxxxxxx",
			want: nil,
		},
		"nil_blunk": {
			sub:  "",
			want: nil,
		},
	}

	for k, tt := range pattern {
		t.Run(k, func(t *testing.T) {
			got := newSubCommand(tt.sub, c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}

}
