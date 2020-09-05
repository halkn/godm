package godm

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type config struct {
	Dirs []string `json:"dirs"`
}

func loadConfig() (*config, error) {
	p := os.Getenv("XDG_CONFIG_HOME")
	if p == "" {
		p = os.Getenv("HOME")
	}
	b, err := ioutil.ReadFile(filepath.Join(p, "godm", "config.json"))
	if err != nil {
		return nil, err
	}

	var cfg config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
