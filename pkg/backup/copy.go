package backup

import (
	"io/fs"
	"path/filepath"

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
