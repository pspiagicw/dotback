package argparse

import (
	"flag"

	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

func ParseArguments(version string) []string {

	Usage := func() {
		help.PrintHelp(version)
	}
	flag.Usage = Usage
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		help.PrintHelp(version)
		goreland.LogFatal("No commands provided!")
	}

	return args
}
