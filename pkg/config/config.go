package config

import (
	"bytes"
	"errors"
	"io/fs"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"github.com/pspiagicw/goreland"
)

type BackupRule struct {
	Location string `toml:"location"`
}
type Config struct {
	StoreDir    string                 `toml:"storeDir"`
	Rules       map[string]*BackupRule `toml:"backup"`
	AfterBackup []string               `toml:"after-backup"`
}

func getConfigPath() string {
	location, err := xdg.ConfigFile("dotback/backup.yml")
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
		goreland.LogError("File %s does not exist, yet. Do create one.", getConfigPath())
		goreland.LogFatal("Run dotback help config for more information.")
	}
}
func GetConfig() *Config {
	assertConfigFile()
	contents := readConfigFile(getConfigPath())
	d := toml.NewDecoder(bytes.NewReader(contents))
	config := new(Config)
	d.Decode(config)
	return config

}
