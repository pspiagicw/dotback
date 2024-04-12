package config

import (
	"flag"
	"fmt"

	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

func parseConfigArgs(opts *argparse.Opts) {
	flag := flag.NewFlagSet("dotback config", flag.ExitOnError)

	flag.Usage = help.HelpConfig

	flag.Parse(opts.Args)
}

func PrintConfig(opts *argparse.Opts) {
	parseConfigArgs(opts)
	fmt.Println("DOTBACK CONFIG")
	config := NewConfig(opts)
	fmt.Printf("Location dir: %s\n", config.StoreDir)
	printAfterBackup(config)
	printRules(config)

}
func printAfterBackup(config *Config) {
	fmt.Println("\nThe after-backup commands:")
	for _, command := range config.AfterBackup {
		goreland.LogExec(command)
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
