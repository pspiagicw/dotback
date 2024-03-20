package help

import (
	"github.com/pspiagicw/pelp"
)

func HelpBackup() {
	pelp.Print("Backup dotfiles")
	pelp.HeaderWithDescription("usage", []string{"dotback [flags] backup <rules>"})

	pelp.Flags("flags", []string{"dry-run", "ignore"}, []string{"Dry run the backup.( Don't execute any disk operations)", "Ignore these backup rules (Backup everything other than these)"})

	pelp.HeaderWithDescription("arguments", []string{"The rules to execute.", "- If empty, backup everything."})
	pelp.Examples("examples", []string{"dotback backup", "dotback backup nvim emacs micro", "dotback --dry-run backup"})
}
func HelpConfig() {
	pelp.Print("Show config info")
	pelp.HeaderWithDescription("usage", []string{"dotback config"})
	pelp.Examples("examples", []string{"dotback config"})
	pelp.HeaderWithDescription("more help", []string{"Use 'dotback help [command]' for more info about a command."})
}
