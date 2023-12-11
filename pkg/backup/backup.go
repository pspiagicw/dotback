package backup

import (
	"path/filepath"

	"github.com/kballard/go-shellquote"
	"github.com/pspiagicw/dotback/pkg/config"
	"github.com/pspiagicw/goreland"
)

func PerformBackup(args []string) {
	configFile := preBackup()

	if len(args) != 0 {
		backupSelective(configFile, args)
	} else {
		backupAll(configFile)

	}
	goreland.LogInfo("Backup complete!")
	runAfterBackup(configFile)
	goreland.LogSuccess("Backup successful!")
}
func preBackup() *config.Config {

	configFile := config.GetConfig()
	storePath := expandHome(configFile.StoreDir)
	goreland.LogInfo("Starting backup at %s", storePath)
	createIfNotExist(storePath)

	return configFile

}
func backupSelective(configFile *config.Config, args []string) {
	for _, name := range args {
		executeRule(configFile, name)
	}
}

func executeRule(configFile *config.Config, name string) {
	rule, exists := configFile.Rules[name]
	if !exists {
		goreland.LogFatal("Could not find [%s] backup rule", name)
	}

	goreland.LogInfo("Backing up [%s]", name)

	src, dest := getPath(configFile, rule)
	performCopy(src, dest)

}
func getPath(configFile *config.Config, rule *config.BackupRule) (string, string) {
	storeDir := expandHome(configFile.StoreDir)
	src := expandHome(rule.Location)
	dest := filepath.Join(storeDir, filepath.Base(src))
	return src, dest
}
func backupAll(configFile *config.Config) {
	for name, _ := range configFile.Rules {
		executeRule(configFile, name)
	}
}
func runAfterBackup(configfile *config.Config) {
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
