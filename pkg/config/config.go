package config

import (
	"bytes"
	"errors"
	"io/fs"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/goreland"
)

type BackupRule struct {
	Location string `toml:"location"`
}
type Config struct {
	StoreDir    string                 `toml:"storeDir"`
	Rules       map[string]*BackupRule `toml:"backup"`
	AfterBackup []string               `toml:"after-backup"`
	Ignore      []string               `toml:"ignore"`
}

func getConfigPath() string {
	location, err := xdg.ConfigFile("dotback/backup.toml")
	if err != nil {
		goreland.LogFatal("Error getting config filepath: %q", err)
	}
	return location
}
func readConfigFile(filepath string) []byte {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		goreland.LogFatal("Error reading config file: %v", err)
	}
	return contents
}
func assertConfigFile() {
	_, err := os.Stat(getConfigPath())
	if errors.Is(err, fs.ErrNotExist) {
		goreland.LogError("The config file '%s' doesn't seem to exist. Create a config to get started", getConfigPath())
		goreland.LogFatal("Run dotback help config for more information.")
	}
}
func NewConfig(opts *argparse.Opts) *Config {

	path := getConfigPath()

	if opts.Config != "" {
		path = opts.Config
	}

	config := newFromFile(path)

	return config

}
func newFromFile(path string) *Config {
	contents := readConfigFile(path)
	d := toml.NewDecoder(bytes.NewReader(contents))
	config := new(Config)
	d.Decode(config)

	checkConfig(config)
	return config
}
func checkConfig(config *Config) {
	if config.StoreDir == "" {
		goreland.LogFatal("StoreDir is not set in the config file")
	}
	if config.Rules == nil {
		goreland.LogFatal("No backup rules found in the config file")
	}
}
