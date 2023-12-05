package backup

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/kballard/go-shellquote"
	homedir "github.com/mitchellh/go-homedir"
	cp "github.com/otiai10/copy"
	"github.com/pspiagicw/dotback/pkg/config"
	"github.com/pspiagicw/goreland"
)

func PerformBackup(args []string) {
	configFile := config.GetConfig()

	storePath := expandHome(configFile.StoreDir)
	goreland.LogInfo("Starting backup at %s", storePath)
	createIfNotExist(storePath)

	startBackUp(configFile)
	runAfterBackup(configFile)

	goreland.LogSuccess("Backup successful!")
}
func runAfterBackup(configfile *config.Config) {
	// fmt.Println(configfile.AfterBackup)
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
func expandHome(filepath string) string {
	path, err := homedir.Expand(filepath)
	if err != nil {
		goreland.LogFatal("Unable to expand home variable: %v", err)
	}
	return path

}
func startBackUp(configFile *config.Config) {
	for name, rule := range configFile.Rules {
		goreland.LogInfo("Backing up [%s]", name)
		storeDir := expandHome(configFile.StoreDir)
		src := expandHome(rule.Location)
		newPath := filepath.Join(storeDir, filepath.Base(src))
		err := cp.Copy(src, newPath, cp.Options{
			Skip: func(srcinfo fs.FileInfo, src string, dest string) (bool, error) {
				if srcinfo.IsDir() && filepath.Base(src) == ".git" {
					return true, nil
				}
				return false, nil
			},
		})
		if err != nil {
			goreland.LogError("Error copying %s: %v", src, err)
		}

	}
}
func createIfNotExist(folder string) {
	if _, err := os.Stat(folder); errors.Is(err, fs.ErrNotExist) {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			goreland.LogFatal("Error creating directory: %s", folder)
		}
	} else if err != nil {
		goreland.LogFatal("Error stating file: %s", err)
	}
}
