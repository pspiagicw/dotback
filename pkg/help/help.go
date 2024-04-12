package help

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/pelp"
)

const EXAMPLE_CONFIG = `
# A folder to store the backup, it will be created if it does not exist.
storeDir = "~/.local/state/backup"

# All commands should be defined by the user.
# It can be left empty or omitted.
after-backup = [
    "scp -r ...",
    "rsync ....",
    "tar -xvzf ..."
]

[backup.nvim]
location = "~/.config/nvim"

# A backup rule has [backup.<rule-name>] format.
# It should contain a 'location' parameter.
[backup.neomutt]
location = "~/.config/neomutt"

# Backup location can also be a file.
[backup.gitconfig]
location = "~/.gitconfig"

`

func PrintVersion(version string) {
	pelp.Version("dotback", version)
}
func printHeader() {
	pelp.Print("Backup dotfiles the simple way!")
	pelp.HeaderWithDescription("usage", []string{"dotback [command] [args]"})
}

func PrintHelp(version string) {
	PrintVersion(version)
	printHeader()
	printCommands()
	printFlags()

	pelp.Examples("examples", []string{"dotback backup", "dotback config"})

	printFooter()
}

func printFlags() {
	flags := []string{"example-config", "config"}
	messages := []string{"Print example config", "Path to the alternate config file."}

	pelp.Flags("flags", flags, messages)
}

func printFooter() {
	pelp.HeaderWithDescription("more help", []string{"Use 'dotback help [command]' for more info about a command."})
}
func printCommands() {
	commands := []string{"backup:", "version:", "config:", "help:"}
	messages := []string{"Backup your dotfiles", "Show version info", "Print the current config", "Show this message"}
	pelp.Aligned("commands", commands, messages)
}
func Handle(args []string, version string) {
	if len(args) == 0 {
		PrintHelp(version)
		return
	}
	cmd := args[0]

	handlers := map[string]func(){
		"backup": HelpBackup,
		"config": HelpConfig,
		"version": func() {
			PrintVersion(version)
		},
	}

	handlerFunc, exists := handlers[cmd]
	if exists {
		handlerFunc()
	} else {
		goreland.LogFatal("No help for command %s found", cmd)
	}

}
func HelpExampleConfig() {
	pelp.HeaderWithDescription("example config", []string{})
	pelp.Print(EXAMPLE_CONFIG)
}
