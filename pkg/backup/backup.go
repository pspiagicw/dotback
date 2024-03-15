package backup

import (
	"flag"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kballard/go-shellquote"
	"github.com/pspiagicw/dotback/pkg/config"
	"github.com/pspiagicw/dotback/pkg/help"
	"github.com/pspiagicw/goreland"
)

type opts struct {
	dryRun bool
	args   []string
	ignore bool
}

func parseBackupOpts(args []string) *opts {
	opts := new(opts)
	flag := flag.NewFlagSet("groom backup", flag.ExitOnError)
	flag.BoolVar(&opts.dryRun, "dry-run", false, "Dry run the backup")
	flag.BoolVar(&opts.ignore, "ignore", false, "Ignore the backup")
	flag.Usage = help.HelpBackup
	flag.Parse(args)

	opts.args = flag.Args()

	return opts
}

func PerformBackup(args []string) {

	opts := parseBackupOpts(args)

	configFile := preBackup(opts)

	if len(args) != 0 {
		backupSelective(configFile, opts)
	} else {
		if opts.ignore {
			goreland.LogFatal("Can't ignore all rules. Please specify the rules to ignore.")
		}
		backupAll(configFile, opts)

	}
	goreland.LogInfo("Backup complete!")
	runAfterBackup(configFile, opts)
	goreland.LogSuccess("Backup successful!")
}

func preBackup(opts *opts) *config.Config {
	configFile := config.GetConfig()
	confirmBackup()
	ensureStorePath(configFile)
	return configFile
}
func ensureStorePath(configFile *config.Config) {
	storePath := expandHome(configFile.StoreDir)
	goreland.LogInfo("Starting backup at %s", storePath)
	createIfNotExist(storePath)
}
func confirmBackup() {
	confirm := false
	prompt := &survey.Confirm{
		Message: "Do you want to start the backup ?",
	}
	survey.AskOne(prompt, &confirm)
	if !confirm {
		goreland.LogFatal("User cancelled the backup!")
	}

}
func backupSelective(configFile *config.Config, opts *opts) {
	if opts.ignore {
		ignoreRules(configFile, opts)
	} else {
		executeSelectiveBackup(configFile, opts)
	}
}
func ignoreRules(configFile *config.Config, opts *opts) {
	ignoredRules := opts.args
	for name, _ := range configFile.Rules {
		if !contains(ignoredRules, name) {
			executeRule(configFile, name, opts)
		} else {
			goreland.LogInfo("Ignoring the [%s] backup", name)
		}
	}
}
func contains(rules []string, name string) bool {
	for _, rule := range rules {
		if rule == name {
			return true
		}
	}
	return false
}
func executeSelectiveBackup(configFile *config.Config, opts *opts) {
	for _, name := range opts.args {
		executeRule(configFile, name, opts)
	}
}

func getRule(name string, configFile *config.Config) *config.BackupRule {
	rule, exists := configFile.Rules[name]

	if !exists {
		goreland.LogFatal("Could not find [%s] backup rule", name)
	}

	return rule

}
func executeRule(configFile *config.Config, name string, opts *opts) {

	rule := getRule(name, configFile)

	goreland.LogInfo("Backing up [%s]", name)

	src, dest := getPath(configFile, rule)

	if !opts.dryRun {
		performCopy(src, dest)
	} else {
		goreland.LogInfo("Move %s -> %s", src, dest)
	}
}
func getPath(configFile *config.Config, rule *config.BackupRule) (string, string) {
	storeDir := expandHome(configFile.StoreDir)
	src := expandHome(rule.Location)
	dest := filepath.Join(storeDir, filepath.Base(src))
	return src, dest
}
func backupAll(configFile *config.Config, opt *opts) {
	for name, _ := range configFile.Rules {
		executeRule(configFile, name, opt)
	}
}
func confirmAfterBackUp() {
	confirm := false
	prompt := survey.Confirm{
		Message: "Run the after-backup procedure ?",
	}
	survey.AskOne(&prompt, &confirm)
	if !confirm {
		goreland.LogFatal("User cancelled the after-backup procedure")
	}
}
func runAfterBackup(configfile *config.Config, opts *opts) {

	if opts.dryRun {
		goreland.LogInfo("DRY RUN: Run after-backup commands.")
		return
	}
	confirmAfterBackUp()

	for _, cmd := range configfile.AfterBackup {

		args, err := shellquote.Split(cmd)

		if err != nil {
			goreland.LogFatal("Failed to parse command '%s':%v", cmd, err)
		}

		goreland.LogExec(cmd)
		err = goreland.Execute(args[0], args[1:], []string{})
		if err != nil {
			goreland.LogFatal("Error executing: %v", err)
		}
	}
}
