package argparse

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/pspiagicw/dotback/backup"
	"github.com/pspiagicw/dotback/config"
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

type ConfigInitCMD struct {
}

func (c *ConfigInitCMD) Run(opts *Opts) error {
	// config.Init(opts.ConfigPath)
	fmt.Println(EXAMPLE_CONFIG)
	return nil
}

type ConfigShowCMD struct {
}

func (c *ConfigShowCMD) Run(opts *Opts) error {
	config.PrintConfig(opts.ConfigPath)
	return nil
}

type ConfigCMD struct {
	Init ConfigInitCMD `cmd:"" help:"Initialize the default config."`
	Show ConfigShowCMD `cmd:"" help:"Initialize the default config."`
}

var CLI struct {
	ConfigPath string `help:"Alternate config file"`

	Backup BackupCMD `cmd:"" help:"Backup your dotfiles!"`
	Config ConfigCMD `cmd:"" help:"Print the current config."`
}

func ParseArguments(version string) *kong.Context {

	ctx := kong.Parse(&CLI)
	ctx.Run(&Opts{CLI.ConfigPath})
	return ctx
}
