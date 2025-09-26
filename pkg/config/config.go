package config

import (
	"bytes"
	"os"

	"github.com/kballard/go-shellquote"
	"github.com/pspiagicw/dotback/pkg/helper"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/goreland"
)

type BackupRule struct {
	Location string `toml:"location"`
	Name     string
}
type Config struct {
	StoreDir    string                 `toml:"storeDir"`
	Rules       map[string]*BackupRule `toml:"backup"`
	AfterBackup []string               `toml:"after-backup"`
	Ignore      []string               `toml:"ignore"`
	Commands    [][]string
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

func NewConfig(opts *argparse.Opts) *Config {

	path := getConfigPath()

	if opts.Config != "" {
		path = opts.Config
	}

	config := newFromFile(path)

	sanitizeConfig(config)

	return config

}

func sanitizeConfig(config *Config) {
	cleanedStorePath := helper.ExpandHome(config.StoreDir)
	helper.CreateIfNotExist(cleanedStorePath)
	config.StoreDir = cleanedStorePath

	for name, rule := range config.Rules {
		cleanedLocation := helper.ExpandHome(rule.Location)
		doesExist := helper.DoesExist(cleanedLocation)

		if !doesExist {
			goreland.LogFatal("Filepath '%s' for rule [%s], doesn't exist", cleanedLocation, name)
		}
		rule.Location = cleanedLocation
		rule.Name = name

	}

	for _, cmd := range config.AfterBackup {
		args, err := shellquote.Split(cmd)

		if err != nil {
			goreland.LogFatal("Failed to parse command '%s': %v", cmd, err)
		}

		config.Commands = append(config.Commands, args)
	}
}

func newFromFile(path string) *Config {
	contents := readConfigFile(path)
	d := toml.NewDecoder(bytes.NewReader(contents))

	config := new(Config)
	d.Decode(config)

	return config
}
