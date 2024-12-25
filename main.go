package main

import (
	"crypto/rand"
	"cursor-reset-go/files"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	// 打开文件
	file, err := files.GetStorageFile()
	if err != nil {
		fmt.Println("打开文件失败", err.Error())
		fmt.Scan()
		return
	}
	defer file.Close()

	// 读取文件内容或创建新的数据
	var data map[string]interface{}

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败", err.Error())
		fmt.Scan()
		return
	}

	if fileInfo.Size() == 0 {
		// 新文件，创建空 map
		data = make(map[string]interface{})
	} else {
		// 现有文件，先备份再读取
		if err := files.BackupFile(file.Name()); err != nil {
			fmt.Println("备份文件失败", err.Error())
			fmt.Scan()
			return
		}

		// 读取 JSON 内容
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&data); err != nil {
			fmt.Println("解析 JSON 失败", err.Error())
			fmt.Scan()
			return
		}
	}

	// 生成新的 ID
	data["telemetry.machineId"] = generateRandomHex(32)
	data["telemetry.macMachineId"] = generateRandomHex(32)
	data["telemetry.devDeviceId"] = uuid.New().String()

	// 准备写入新内容
	if err := file.Truncate(0); err != nil {
		fmt.Println("截断文件失败", err.Error())
		fmt.Scan()
		return
	}
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("重置文件指针失败", err.Error())
		fmt.Scan()
		return
	}

	// 写入更新后的数据
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		fmt.Println("写入 JSON 失败", err.Error())
		fmt.Scan()
		return
	}

	// 显示结果
	displayResult(data)

	fmt.Println("\n按任意键继续...")
	fmt.Scan()
}

// 显示结果的辅助函数
func displayResult(data map[string]interface{}) {
	fmt.Println("🎉 Device IDs have been successfully reset. The new device IDs are: \n")

	displayData := map[string]string{
		"machineId":    data["telemetry.machineId"].(string),
		"macMachineId": data["telemetry.macMachineId"].(string),
		"devDeviceId":  data["telemetry.devDeviceId"].(string),
	}

	prettyJSON, err := json.MarshalIndent(displayData, "", "  ")
	if err != nil {
		fmt.Println("格式化 JSON 失败", err.Error())
		return
	}
	fmt.Println(string(prettyJSON))
}

func generateRandomHex(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
