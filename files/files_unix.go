//go:build unix && !linux
// +build unix,!linux

package files

import (
	"fmt"
	"os"
	"time"
)

func BackupFile(filePath string) error {
	// 检查源文件是否存在
	srcInfo, err := os.Stat(filePath)
	if err != nil {
		return nil // 如果文件不存在，直接返回
	}

	// 获取原文件的权限模式
	srcMode := srcInfo.Mode()

	// 读取源文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}

	// 创建带时间戳的备份文件名
	timestamp := time.Now().Format("20060102_150405")
	backupPath := fmt.Sprintf("%s.backup_%s", filePath, timestamp)

	// 使用原文件的权限模式写入备份文件
	if err := os.WriteFile(backupPath, data, srcMode.Perm()); err != nil {
		return fmt.Errorf("创建备份文件失败: %v", err)
	}

	return nil
}

func getFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "Library", "Application Support", "Cursor", "User", "globalStorage", "storage.json")
}
