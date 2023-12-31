package main

import (
	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/handle"
)

var VERSION string

func main() {
	args := argparse.ParseArguments(VERSION)
	handler.HandleArgs(args, VERSION)
}
