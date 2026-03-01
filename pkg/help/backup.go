package help

import (
	"github.com/pspiagicw/pelp"
)

func HelpBackup() {
	pelp.Print("Backup dotfiles")
	pelp.HeaderWithDescription("usage", []string{"dotback backup [flags] [rules...]"})

	pelp.Flags(
		"flags",
		[]string{"dry-run", "ignore", "yes", "non-interactive", "no-after-backup"},
		[]string{
			"Preview backup operations without writing files",
			"Treat provided rule names as exclusions",
			"Skip confirmation prompts",
			"Run without interactive prompts",
			"Skip running after-backup commands",
		},
	)

	pelp.HeaderWithDescription("arguments", []string{"The rules to execute.", "- If empty, backup everything."})
	pelp.Examples("examples", []string{"dotback backup", "dotback backup nvim emacs micro", "dotback backup --dry-run", "dotback backup --yes --no-after-backup"})
}
func HelpConfig() {
	pelp.Print("Show resolved config information")
	pelp.HeaderWithDescription("usage", []string{"dotback config"})
	pelp.Examples("examples", []string{"dotback config"})
	pelp.HeaderWithDescription("more help", []string{"Use 'dotback help <command>' for more info about a command."})
}

func HelpList() {
	pelp.Print("List configured backup rule names")
	pelp.HeaderWithDescription("usage", []string{"dotback list"})
	pelp.Examples("examples", []string{"dotback list", "dotback --config /abs/path/backup.toml list"})
	pelp.HeaderWithDescription("more help", []string{"Use 'dotback help <command>' for more info about a command."})
}
