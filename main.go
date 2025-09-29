package main

import (
	"github.com/pspiagicw/dotback/pkg/argparse"
)

var VERSION string = "unversioned"

func main() {
	// opts := argparse.ParseArguments(VERSION)
	// handler.Handle(opts)
	argparse.ParseArguments(VERSION)
	// handler.Handle(ctx)
}
