package handler

import (
	"github.com/pspiagicw/dotback/pkg/helper"
	"github.com/pspiagicw/goreland"
)

func HandleArgs(args []string, version string) {
	cmd := args[0]

	PrintVersion := func([]string) {
		goreland.LogInfo("dotback, version: %s", version)
	}

	handlers := map[string]func([]string){
		"version": PrintVersion,
		"backup":  notImplemented,
	}

	handler, exists := handlers[cmd]
	if !exists {
		helper.PrintHelp(version)
	} else {
		handler(args)
	}
}
func notImplemented(args []string) {
	goreland.LogError("This feature is not implemented yet!")
}
