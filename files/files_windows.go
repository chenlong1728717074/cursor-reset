//go:build windows
// +build windows

package files

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func GetStorageFile() (*os.File, error) {
	var file *os.File
	var err error
	appData := os.Getenv("APPDATA")

	filePath := filepath.Join(appData, "Cursor", "User", "globalStorage", "storage.json")

	file, err = os.OpenFile(
		filePath,
		os.O_RDWR|os.O_CREATE,
		0666,
	)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func BackupFile(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		timestamp := time.Now().Format("20060102_150405")
		backupPath := fmt.Sprintf("%s.backup_%s", filePath, timestamp)
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		return os.WriteFile(backupPath, data, 0644)
	}
	return nil
}
