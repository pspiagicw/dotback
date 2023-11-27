package main

import (
	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/handler"
)

var VERSION string

func main() {
	args := argparse.ParseArguments(VERSION)
	handler.HandleArgs(args, VERSION)
}
