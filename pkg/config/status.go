package config

import (
	"fmt"

	"github.com/pspiagicw/goreland"
)

func PrintConfig(args []string) {
	fmt.Println("DOTBACK CONFIG")
	config := GetConfig()
	fmt.Printf("Location dir: %s\n", config.StoreDir)
	printAfterBackup(config)
	printRules(config)

}
func printAfterBackup(config *Config) {
	fmt.Println("\nThe after-backup commands:\n")
	for _, command := range config.AfterBackup {
		goreland.LogExec(command)
	}
}
func printRules(config *Config) {
	fmt.Println("\nConfigured backup rules:\n")
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
