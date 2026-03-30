package main

import (
	"github.com/pspiagicw/dotback/argparse"
)

var VERSION string = "unversioned"

func main() {
	argparse.ParseArguments(VERSION)
}
