package argparse

import (
	"flag"

	"github.com/pspiagicw/dotback/pkg/helper"
)

func ParseArguments(version string) []string {

	Usage := func() {
		helper.PrintHelp(version)
	}
	flag.Usage = Usage
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		helper.PrintHelp(version)
	}

	return args
}
