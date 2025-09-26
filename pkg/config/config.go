package config

import (
	"bytes"
	"os"
	"time"

	"github.com/kballard/go-shellquote"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/pspiagicw/demp"
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
	cleanedStorePath := goreland.ExpandHome(config.StoreDir)
	goreland.CreateIfNotExist(cleanedStorePath)
	config.StoreDir = cleanedStorePath

	for name, rule := range config.Rules {
		cleanedLocation := goreland.ExpandHome(rule.Location)
		doesExist := goreland.DoesExist(cleanedLocation)

		if !doesExist {
			goreland.LogFatal("Filepath '%s' for rule [%s], doesn't exist", cleanedLocation, name)
		}
		rule.Location = cleanedLocation
		rule.Name = name

	}

	for _, cmd := range config.AfterBackup {
		variables := generateVariables(config)
		templatedCmd := demp.ResolveTemplate(cmd, variables)
		args, err := shellquote.Split(templatedCmd)

		if err != nil {
			goreland.LogFatal("Failed to parse command '%s': %v", cmd, err)
		}

		config.Commands = append(config.Commands, args)
	}
}
func generateVariables(config *Config) map[string]string {
	vars := map[string]string{
		"STOREDIR": config.StoreDir,
		"DATE":     time.Now().Local().Format(time.DateOnly),
		"TIME":     time.Now().Local().Format(time.TimeOnly),
	}

	return vars
}

func newFromFile(path string) *Config {
	contents := readConfigFile(path)
	d := toml.NewDecoder(bytes.NewReader(contents))

	config := new(Config)
	d.Decode(config)

	return config
}
