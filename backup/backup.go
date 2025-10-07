package backup

import (
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pspiagicw/dotback/config"
	"github.com/pspiagicw/goreland"
)

type opts struct {
	dryRun bool
	ignore bool
	args   []string
}

func Backup(configPath string, dryRun bool, ignore bool, rules []string) {

	configFile := config.NewConfig(configPath)

	o := &opts{
		dryRun: dryRun,
		ignore: ignore,
		args:   rules,
	}

	startBackup(configFile, o)
	postBackup(configFile, o)
}
func logRules(rules []string, o *opts) {

	if len(rules) == 0 && len(o.args) != 1 {
		goreland.LogFatal("No rules to execute, check command or arguments.")
	}

	if o.ignore {
		goreland.LogInfo("Ignoring rules: %s", o.args)
	}

	goreland.LogInfo("Executing rules: %s", rules)
	goreland.Confirm("Do you want to start the backup ?", "User cancelled the backup")
}

func startBackup(configFile *config.Config, o *opts) {
	rulesToExecute := filterRules(configFile, o)

	logRules(rulesToExecute, o)

	for _, rule := range rulesToExecute {
		executeRule(configFile, rule, o)
	}
}

func filterRules(configFile *config.Config, o *opts) []string {

	keys := mapset.NewSetFromMapKeys(configFile.Rules)
	args := mapset.NewSet(o.args...)

	// If no args given, execute everything (only when ignore not given)
	if !o.ignore && len(o.args) == 0 {
		args = keys
	}

	var results mapset.Set[string]

	if o.ignore {
		results = keys.Difference(args)
	} else {
		results = keys.Intersect(args)
	}

	return results.ToSlice()
}

func postBackup(configFile *config.Config, o *opts) {
	goreland.LogInfo("Backup complete!")
	goreland.Confirm("Run the after-backup procedure ?", "User cancelled the after-backup procedure.")

	runAfterBackup(configFile, o)

	goreland.LogSuccess("Backup successful!")
}

func executeRule(configFile *config.Config, name string, o *opts) {

	rule := configFile.Rules[name]

	goreland.LogInfo("Backing up [%s]", name)

	src := rule.Location
	dest := filepath.Join(configFile.StoreDir, filepath.Base(src))

	if !o.dryRun {
		goreland.CopyIgnoreGit(src, dest, configFile.Ignore)
	} else {
		goreland.LogInfo("Move %s -> %s", src, dest)
	}
}

func runAfterBackup(configFile *config.Config, opts *opts) {

	for i := range configFile.AfterBackup {
		cmd := configFile.AfterBackup[i]
		command := configFile.Commands[i]

		goreland.LogExecSimple(cmd)

		if opts.dryRun {
			continue
		}

		err := goreland.ExecuteDir(command[0], command[1:], []string{}, configFile.StoreDir)
		if err != nil {
			goreland.LogFatal("Error executing: %v", err)
		}
	}
}
