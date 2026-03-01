package main

import (
	"github.com/pspiagicw/dotback/argparse"
)

var VERSION string = "unversioned"

func main() {
	// opts := argparse.ParseArguments(VERSION)
	// handler.Handle(opts)
	argparse.ParseArguments(VERSION)
	// handler.Handle(ctx)
}
