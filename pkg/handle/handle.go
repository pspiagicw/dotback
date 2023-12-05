package handler

import (
	"github.com/pspiagicw/dotback/pkg/backup"
	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

func HandleArgs(args []string, version string) {
	cmd := args[0]

	handlers := map[string]func([]string){
		"version": func([]string) {
			help.PrintVersion(version)
		},
		"backup": backup.PerformBackup,
		"help":   notImplemented,
	}

	handler, exists := handlers[cmd]
	if exists {
		handler(args)
	} else {
		help.PrintHelp(version)
		goreland.LogFatal("No command named %s", cmd)
	}
}
func notImplemented(args []string) {
	goreland.LogError("This feature is not implemented yet!")
}
