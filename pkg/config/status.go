package config

import (
	"fmt"

	"github.com/pspiagicw/goreland"
)

func PrintConfig(configPath string) {
	config := NewConfig(configPath)
	fmt.Printf("Location dir: %s\n", config.StoreDir)
	printAfterBackup(config)
	printRules(config)
}

func printAfterBackup(config *Config) {
	fmt.Println("\nThe after-backup commands:")
	for _, command := range config.AfterBackup {
		goreland.LogExecSimple(command)
	}
}
func printRules(config *Config) {
	fmt.Println("\nConfigured backup rules:")

	headers := []string{
		"Name",
		"Location",
	}

	rows := [][]string{}
	for name, rule := range config.Rules {
		rows = append(rows, []string{name, rule.Location})
	}

	goreland.LogTable(headers, rows)

}
