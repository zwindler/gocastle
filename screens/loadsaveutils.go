package screens

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

func getBaseDirectory() (fyne.ListableURI, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	absoluteDirPath := filepath.Dir(executablePath)
	return storage.ListerForURI(storage.NewFileURI(absoluteDirPath))
}
