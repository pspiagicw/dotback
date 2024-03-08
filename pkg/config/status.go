package config

import (
	"flag"
	"fmt"

	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

func parseConfigArgs(args []string) {
	flag := flag.NewFlagSet("dotback config", flag.ExitOnError)

	flag.Usage = help.HelpConfig

	flag.Parse(args)
}

func PrintConfig(args []string) {
	parseConfigArgs(args)
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
