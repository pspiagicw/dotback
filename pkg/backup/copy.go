package backup

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	cp "github.com/otiai10/copy"
	"github.com/pspiagicw/goreland"
)

func performCopy(src string, dest string, ignore []string) {
	err := cp.Copy(src, dest, cp.Options{
		Skip: generateSkipFunc(ignore),
	})
	if err != nil {
		goreland.LogError("Error copying %s: %v", src, err)
	}
}
func generateSkipFunc(ignore []string) func(srcinfo fs.FileInfo, src string, dest string) (bool, error) {
	return func(srcinfo fs.FileInfo, src string, dest string) (bool, error) {
		if srcinfo.IsDir() && filepath.Base(src) == ".git" {
			return true, nil
		}
		for _, pattern := range ignore {
			match, err := filepath.Match(pattern, src)
			if err != nil {
				return false, err
			}
			if match {
				return true, nil
			}
		}
		return false, nil
	}
}

func SkipFunc(srcinfo fs.FileInfo, src string, dest string) (bool, error) {
	if srcinfo.IsDir() && filepath.Base(src) == ".git" {
		return true, nil
	}
	return false, nil
}
func expandHome(filepath string) string {
	path, err := homedir.Expand(filepath)
	if err != nil {
		goreland.LogFatal("Unable to expand home variable: %v", err)
	}
	return path
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
