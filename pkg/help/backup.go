package help

import "fmt"

func HelpBackup() {
	fmt.Println("Backup dotfiles")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("   dotback [flags] backup <rules>")
	fmt.Println()
	fmt.Println("FLAGS")
	flags := `
--dry-run:`
	description := `
Dry run the backup.( Don't execute any disk operations)`
	printAligned(flags, description)
	fmt.Println("ARGUMENTS")
	fmt.Println("   The rules to execute.")
	fmt.Println("   - If empty, backup everything.")
	fmt.Println()
	fmt.Println("EXAMPLES")
	fmt.Println("   $ dotback backup")
	fmt.Println("   $ dotback backup nvim emacs micro")
	fmt.Println("   $ dotback --dry-run backup")
}
func HelpConfig() {
	fmt.Println("Show config info")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("   dotback config")
	fmt.Println()
	fmt.Println("EXAMPLES")
	fmt.Println("   $ dotback config")
}
