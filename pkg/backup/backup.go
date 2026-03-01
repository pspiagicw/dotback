package backup

import (
	"fmt"
	"flag"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kballard/go-shellquote"
	"github.com/pspiagicw/dotback/pkg/argparse"
	"github.com/pspiagicw/dotback/pkg/config"
	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

func parseBackupOpts(opts *argparse.Opts) {

	flag := flag.NewFlagSet("groom backup", flag.ExitOnError)

	flag.BoolVar(&opts.DryRun, "dry-run", false, "Dry run the backup")
	flag.BoolVar(&opts.Ignore, "ignore", false, "Ignore the backup")
	flag.BoolVar(&opts.Yes, "yes", false, "Skip confirmation prompts")
	flag.BoolVar(&opts.NonInteractive, "non-interactive", false, "Run without interactive prompts")
	flag.BoolVar(&opts.NoAfterBackup, "no-after-backup", false, "Skip running after-backup commands")

	flag.Usage = help.HelpBackup
	flag.Parse(opts.Args)

	opts.Args = flag.Args()
}

func Backup(opts *argparse.Opts) {

	parseBackupOpts(opts)
	configFile := getConfig(opts)

	failures := executeBackup(configFile, opts)
	goreland.LogInfo("Backup complete!")
	failures = append(failures, runAfterBackup(configFile, opts)...)
	if len(failures) != 0 {
		for _, failure := range failures {
			goreland.LogError("%s", failure)
		}
		goreland.LogFatal("Backup finished with %d failure(s)", len(failures))
	}
	goreland.LogSuccess("Backup successful!")
}

func executeBackup(configFile *config.Config, opts *argparse.Opts) []string {
	if len(opts.Args) != 0 {
		return backupSelective(configFile, opts)
	}
	return backupAll(configFile, opts)
}

func getConfig(opts *argparse.Opts) *config.Config {
	configFile := config.NewConfig(opts)
	confirmBackup(opts)
	ensureStorePath(configFile)
	return configFile
}
func ensureStorePath(configFile *config.Config) {
	storePath := expandHome(configFile.StoreDir)
	goreland.LogInfo("Starting backup at %s", storePath)
	createIfNotExist(storePath)
}
func isNonInteractive(opts *argparse.Opts) bool {
	return opts.Yes || opts.NonInteractive
}
func confirmBackup(opts *argparse.Opts) {
	if isNonInteractive(opts) {
		goreland.LogInfo("Skipping backup confirmation prompt")
		return
	}
	confirm := false
	prompt := &survey.Confirm{
		Message: "Do you want to start the backup ?",
	}
	survey.AskOne(prompt, &confirm)
	if !confirm {
		goreland.LogFatal("User cancelled the backup!")
	}

}
func backupSelective(configFile *config.Config, opts *argparse.Opts) []string {
	if opts.Ignore {
		return ignoreRules(configFile, opts)
	}
	return executeSelectiveBackup(configFile, opts)
}
func ignoreRules(configFile *config.Config, opts *argparse.Opts) []string {
	ignoredRules := opts.Args
	failures := []string{}
	for name, _ := range configFile.Rules {
		if !contains(ignoredRules, name) {
			if err := executeRule(configFile, name, opts); err != nil {
				failures = append(failures, err.Error())
			}
		} else {
			goreland.LogInfo("Ignoring the [%s] backup", name)
		}
	}
	return failures
}
func contains(rules []string, name string) bool {
	for _, rule := range rules {
		if rule == name {
			return true
		}
	}
	return false
}
func executeSelectiveBackup(configFile *config.Config, opts *argparse.Opts) []string {
	failures := []string{}
	for _, name := range opts.Args {
		if err := executeRule(configFile, name, opts); err != nil {
			failures = append(failures, err.Error())
		}
	}
	return failures
}

func getRule(name string, configFile *config.Config) *config.BackupRule {
	rule, exists := configFile.Rules[name]

	if !exists {
		goreland.LogFatal("Could not find [%s] backup rule", name)
	}

	return rule

}
func executeRule(configFile *config.Config, name string, opts *argparse.Opts) error {

	rule := getRule(name, configFile)

	goreland.LogInfo("Backing up [%s]", name)

	src, dest := getPath(configFile, rule)

	if !opts.DryRun {
		if err := performCopy(src, dest, configFile.Ignore); err != nil {
			return fmt.Errorf("failed rule [%s]: copy %s -> %s: %w", name, src, dest, err)
		}
	} else {
		goreland.LogInfo("Move %s -> %s", src, dest)
	}
	return nil
}
func getPath(configFile *config.Config, rule *config.BackupRule) (string, string) {
	storeDir := expandHome(configFile.StoreDir)
	src := expandHome(rule.Location)
	dest := filepath.Join(storeDir, filepath.Base(src))
	return src, dest
}
func backupAll(configFile *config.Config, opt *argparse.Opts) []string {
	if opt.Ignore {
		goreland.LogFatal("Can't ignore all rules")
	}

	failures := []string{}
	for name, _ := range configFile.Rules {
		if err := executeRule(configFile, name, opt); err != nil {
			failures = append(failures, err.Error())
		}
	}
	return failures
}
func confirmAfterBackUp(opts *argparse.Opts) {
	if isNonInteractive(opts) {
		goreland.LogInfo("Skipping after-backup confirmation prompt")
		return
	}
	confirm := false
	prompt := survey.Confirm{
		Message: "Run the after-backup procedure ?",
	}
	survey.AskOne(&prompt, &confirm)
	if !confirm {
		goreland.LogFatal("User cancelled the after-backup procedure")
	}
}
func runAfterBackup(configfile *config.Config, opts *argparse.Opts) []string {

	if opts.DryRun {
		goreland.LogInfo("DRY RUN: Run after-backup commands.")
		return nil
	}
	if opts.NoAfterBackup {
		goreland.LogInfo("Skipping after-backup commands (--no-after-backup)")
		return nil
	}
	if len(configfile.AfterBackup) == 0 {
		return nil
	}

	confirmAfterBackUp(opts)
	failures := []string{}

	for _, cmd := range configfile.AfterBackup {

		args, err := shellquote.Split(cmd)

		if err != nil {
			failures = append(failures, fmt.Sprintf("failed to parse after-backup command '%s': %v", cmd, err))
			continue
		}
		if len(args) == 0 {
			failures = append(failures, "empty after-backup command")
			continue
		}

		goreland.LogExec(cmd)
		err = goreland.Execute(args[0], args[1:], []string{})
		if err != nil {
			failures = append(failures, fmt.Sprintf("failed to execute after-backup command '%s': %v", cmd, err))
		}
	}
	return failures
}
