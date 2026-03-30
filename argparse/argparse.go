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
	Version    string
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

type ConfigDefaultCMD struct {
}

func (c *ConfigDefaultCMD) Run(opts *Opts) error {
	// config.Init(opts.ConfigPath)
	fmt.Print(EXAMPLE_CONFIG)
	return nil
}

type ConfigShowCMD struct {
}

func (c *ConfigShowCMD) Run(opts *Opts) error {
	config.PrintConfig(opts.ConfigPath)
	return nil
}

type ConfigCMD struct {
	Default ConfigDefaultCMD `cmd:"" help:"Initialize the default config."`
	Show    ConfigShowCMD    `cmd:"" help:"Initialize the default config."`
}

type VersionCMD struct {
}

func (v *VersionCMD) Run(opts *Opts) error {
	fmt.Printf("dotback version: %s\n", opts.Version)
	return nil
}

var CLI struct {
	ConfigPath string `help:"Alternate config file"`

	Backup  BackupCMD  `cmd:"" help:"Backup your dotfiles!"`
	Config  ConfigCMD  `cmd:"" help:"Print the current config."`
	Version VersionCMD `cmd:"" help:"Print version info"`
}

func ParseArguments(version string) *kong.Context {

	ctx := kong.Parse(&CLI)
	ctx.Run(&Opts{CLI.ConfigPath, version})
	return ctx
}
