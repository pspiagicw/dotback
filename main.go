package main

import (
	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/handle"
)

var VERSION string = "unversioned"

func main() {
	opts := argparse.ParseArguments(VERSION)
	handler.Handle(opts)
}
