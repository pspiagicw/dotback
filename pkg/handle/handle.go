package handler

import (
	"os"

	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/backup"
	"github.com/pspiagicw/dotback/pkg/config"
	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

func HandleArgs(opts *argparse.Opts) {

	checkExampleConfig(opts)
	checkArgLen(opts)

	cmd := opts.Args[0]

	handlers := map[string]func(*argparse.Opts){
		"version": func(*argparse.Opts) {
			help.PrintVersion(opts.Version)
		},
		"backup": backup.PerformBackup,
		"config": config.PrintConfig,
		"help": func(opts *argparse.Opts) {
			help.HelpArgs(opts.Args, opts.Version)
		},
	}

	handler, exists := handlers[cmd]

	opts.Args = opts.Args[1:]

	if exists {
		handler(opts)
	} else {
		help.PrintHelp(opts.Version)
		goreland.LogFatal("No command named %s", cmd)
	}
}
func checkExampleConfig(opts *argparse.Opts) {
	if opts.ExampleConfig {
		help.HelpExampleConfig()
		os.Exit(0)
	}
}
func checkArgLen(opts *argparse.Opts) {
	if len(opts.Args) == 0 {
		help.PrintHelp(opts.Version)
		goreland.LogFatal("No command provided")
	}

}
