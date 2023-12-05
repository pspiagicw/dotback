package help

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
)

func PrintVersion(version string) {
	fmt.Printf("dotback version %s\n", version)
}

func PrintHelp(version string) {
	goreland.LogInfo("dotback version: %s", version)
	fmt.Println("Backup dotfiles the simple way!")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  dotback [command] [args]")
	fmt.Println()
	fmt.Println("COMMANDS")
	commands := `
backup:
version:
help:`
	messages := `
Backup your dotfiles
Show version info
Show this message`

	commandCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(commands).MarginLeft(2).String()
	messageCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(messages).MarginLeft(5).String()

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, commandCol, messageCol))

	fmt.Println()
	fmt.Println("MORE HELP")
	fmt.Println("  Use 'dotback help [command]' for more info about a command.")

	fmt.Println()
	fmt.Println("EXAMPLES")
	fmt.Println("  $ dotback backup")
	fmt.Println()
}
