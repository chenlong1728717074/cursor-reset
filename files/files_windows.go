//go:build windows
// +build windows

package files

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

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

func getFilePath() string {
	appData := os.Getenv("APPDATA")
	return filepath.Join(appData, "Cursor", "User", "globalStorage", "storage.json")
}
