package file

import "os"

func CreateDir(folder string) error {
	if IsDirExisted(folder) {
		return nil
	}
	return os.MkdirAll(folder, os.ModePerm)
}

func IsDirExisted(path string) bool {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if stat.IsDir() {
		return true
	}
	return false
}
