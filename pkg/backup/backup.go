package backup

import (
	"flag"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/config"
	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

func parseBackupOpts(opts *argparse.Opts) {

	flag := flag.NewFlagSet("groom backup", flag.ExitOnError)

	flag.BoolVar(&opts.DryRun, "dry-run", false, "Dry run the backup")
	flag.BoolVar(&opts.Ignore, "ignore", false, "Ignore the backup")

	flag.Usage = help.HelpBackup
	flag.Parse(opts.Args)

	opts.Args = flag.Args()
}

func Backup(opts *argparse.Opts) {

	parseBackupOpts(opts)
	configFile := config.NewConfig(opts)
	goreland.Confirm("Do you want to start the backup ?", "User cancelled the backup")

	startBackup(configFile, opts)
	postBackup(configFile, opts)
}

func startBackup(configFile *config.Config, opts *argparse.Opts) {
	rulesToExecute := filterRules(configFile, opts)

	for _, rule := range rulesToExecute {
		executeRule(configFile, rule, opts)
	}
}

func filterRules(configFile *config.Config, opts *argparse.Opts) []string {

	keys := mapset.NewSetFromMapKeys(configFile.Rules)
	args := mapset.NewSet(opts.Args...)

	// If no args given, execute everything (only when ignore not given)
	if !opts.Ignore && len(opts.Args) == 0 {
		args = keys
	}

	var results mapset.Set[string]

	if opts.Ignore {
		results = keys.Difference(args)
	} else {
		results = keys.Intersect(args)
	}

	return results.ToSlice()
}

func postBackup(configFile *config.Config, opts *argparse.Opts) {
	goreland.LogInfo("Backup complete!")
	goreland.Confirm("Run the after-backup procedure ?", "User cancelled the after-backup procedure.")

	runAfterBackup(configFile, opts)

	goreland.LogSuccess("Backup successful!")
}

func executeRule(configFile *config.Config, name string, opts *argparse.Opts) {

	rule := configFile.Rules[name]

	goreland.LogInfo("Backing up [%s]", name)

	src := rule.Location
	dest := filepath.Join(configFile.StoreDir, filepath.Base(src))

	if !opts.DryRun {
		goreland.CopyIgnoreGit(src, dest, configFile.Ignore)
	} else {
		goreland.LogInfo("Move %s -> %s", src, dest)
	}
}

func runAfterBackup(configFile *config.Config, opts *argparse.Opts) {

	for i := range configFile.AfterBackup {
		cmd := configFile.AfterBackup[i]
		command := configFile.Commands[i]

		goreland.LogExecSimple(cmd)

		if opts.DryRun {
			continue
		}

		err := goreland.ExecuteDir(command[0], command[1:], []string{}, configFile.StoreDir)
		if err != nil {
			goreland.LogFatal("Error executing: %v", err)
		}
	}
}
