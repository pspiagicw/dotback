package helper

import (
	"errors"
	"io/fs"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/pspiagicw/goreland"
)

func CreateIfNotExist(folder string) {
	if _, err := os.Stat(folder); errors.Is(err, fs.ErrNotExist) {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			goreland.LogFatal("Error creating directory: %s", folder)
		}
	} else if err != nil {
		goreland.LogFatal("Error stating file: %s", err)
	}
}

func ExpandHome(filepath string) string {
	path, err := homedir.Expand(filepath)
	if err != nil {
		goreland.LogFatal("Unable to expand home variable: %v", err)
	}
	return path
}
func DoesExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}
func Confirm(message string, failedmessage string) {
	confirm := false
	prompt := &survey.Confirm{
		Message: message,
	}
	survey.AskOne(prompt, &confirm)
	if !confirm {
		goreland.LogFatal(failedmessage)
	}
}
