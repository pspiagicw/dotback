package handler

import (
	"os"

	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/backup"
	"github.com/pspiagicw/dotback/pkg/config"
	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

func Handle(opts *argparse.Opts) {

	if opts.ExampleConfig {
		printExampleConfig()
	} else if len(opts.Args) == 0 {
		noCommandProvided(opts)
	}

	executeCommand(opts)
}

func executeCommand(opts *argparse.Opts) {
	cmd := opts.Args[0]
	opts.Args = opts.Args[1:]

	switch cmd {
	case "version":
		help.Version(opts.Version)
	case "backup":
		backup.Backup(opts)
	case "config":
		config.PrintConfig(opts)
	case "help":
		help.Handle(opts.Args, opts.Version)
	default:
		help.PrintHelp(opts.Version)
		goreland.LogFatal("No command named %s", cmd)
	}
}

func printExampleConfig() {
	help.HelpExampleConfig()
	os.Exit(0)
}

func noCommandProvided(opts *argparse.Opts) {
	help.PrintHelp(opts.Version)
	goreland.LogFatal("No command provided")
}
