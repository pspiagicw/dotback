package handler

import (
	"os"

	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/backup"
	"github.com/pspiagicw/dotback/pkg/config"
	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

var handlers map[string]func(*argparse.Opts) = map[string]func(*argparse.Opts){
	"version": func(opts *argparse.Opts) {
		help.PrintVersion(opts.Version)
	},
	"backup": backup.Backup,
	"config": config.PrintConfig,
	"help": func(opts *argparse.Opts) {
		help.Handle(opts.Args, opts.Version)
	},
}

func Handle(opts *argparse.Opts) {

	checkExampleConfig(opts)
	checkArgLen(opts)
	execCmd(opts)
}

func execCmd(opts *argparse.Opts) {
	cmd := opts.Args[0]
	opts.Args = opts.Args[1:]

	handler := handlers[cmd]

	if handler != nil {
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
