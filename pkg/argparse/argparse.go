package argparse

import (
	"flag"

	"github.com/pspiagicw/dotback/pkg/help"
)

type Opts struct {
	ExampleConfig bool
	Args          []string
	Version       string

	// Config file path
	Config string

	// Backup options
	DryRun bool
	Ignore bool
}

func ParseArguments(version string) *Opts {

	opts := new(Opts)

	Usage := func() {
		help.PrintHelp(version)
	}
	flag.BoolVar(&opts.ExampleConfig, "example-config", false, "Print example config.")
	flag.StringVar(&opts.Config, "config", "", "Path to the alternate config file.")
	flag.Usage = Usage
	flag.Parse()

	opts.Args = flag.Args()
	opts.Version = version

	return opts
}
