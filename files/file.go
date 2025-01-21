package files

import "os"

func GetStorageFile() (*os.File, error) {
	var file *os.File
	var err error

	var filePath string

	// 获取存储路径
	filePath = getFilePath()

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
