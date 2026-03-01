package config

import (
	"fmt"
	"sort"

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
	sort.Slice(rows, func(i, j int) bool {
		return rows[i][0] < rows[j][0]
	})

	goreland.LogTable(headers, rows)

}

func ListRules(configPath string) {
	config := NewConfig(configPath)

	names := []string{}
	for name := range config.Rules {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name)
	}
}
