package argparse

import (
	"github.com/alecthomas/kong"
	"github.com/pspiagicw/dotback/pkg/backup"
	"github.com/pspiagicw/dotback/pkg/config"
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

type Opts struct {
	ConfigPath string
}

type BackupCMD struct {
	DryRun bool     `help:"Dry run the backup."`
	Ignore bool     `help:"Ignore these backup rules."`
	Rules  []string `arg:"" optional:"" help:"Rules to backup/ignore"`
}

func (b *BackupCMD) Run(opts *Opts) error {
	backup.Backup(opts.ConfigPath, b.DryRun, b.Ignore, b.Rules)
	return nil
}

type ConfigCMD struct {
}

func (c *ConfigCMD) Run(opts *Opts) error {
	config.PrintConfig(opts.ConfigPath)
	return nil
}

type ExampleConfigCMD struct {
}

func (e *ExampleConfigCMD) Run(opts *Opts) error {
	pelp.Print(EXAMPLE_CONFIG)
	return nil
}

var CLI struct {
	ConfigPath string `help:"Alternate config file"`

	Backup        BackupCMD        `cmd:"" help:"Backup your dotfiles!"`
	Config        ConfigCMD        `cmd:"" help:"Print the current config."`
	ExampleConfig ExampleConfigCMD `cmd:"" help:"Print the example config."`
}

func ParseArguments(version string) *kong.Context {

	ctx := kong.Parse(&CLI)
	ctx.Run(&Opts{CLI.ConfigPath})
	return ctx
}
