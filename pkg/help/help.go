package help

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
)

func PrintVersion(version string) {
	fmt.Printf("dotback version %s\n", version)
}
func printHeader() {
	fmt.Println("Backup dotfiles the simple way!")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  dotback [command] [args]")
	fmt.Println()

}

func PrintHelp(version string) {
	PrintVersion(version)
	printHeader()
	printCommands()

	fmt.Println("EXAMPLES")
	fmt.Println("  $ dotback backup")
	fmt.Println()

	printFooter()
}

func printFooter() {
	fmt.Println("MORE HELP")
	fmt.Println("  Use 'dotback help [command]' for more info about a command.")
}
func printCommands() {
	fmt.Println("COMMANDS")
	commands := `
backup:
version:
config:
help:`
	messages := `
Backup your dotfiles
Show version info
Print the current config
Show this message`
	printAligned(commands, messages)
	fmt.Println()
}
func printAligned(left, right string) {
	leftCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(left).MarginLeft(2).String()
	rightCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(right).MarginLeft(5).String()

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, leftCol, rightCol))

	fmt.Println()

}
func HelpArgs(args []string, version string) {
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
