//go:build linux
// +build linux

package files

import (
	"fmt"
	"os"
	"time"
)

func BackupFile(filePath string) error {
	srcInfo, err := os.Stat(filePath)
	if err != nil {
		return nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	backupPath := fmt.Sprintf("%s.backup_%s", filePath, timestamp)

	// Linux 特定实现：保留所有权限位
	if err := os.WriteFile(backupPath, data, srcInfo.Mode()); err != nil {
		return fmt.Errorf("创建备份文件失败: %v", err)
	}

	return nil
}

func getFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", "Cursor", "User", "globalStorage", "storage.json")
}
