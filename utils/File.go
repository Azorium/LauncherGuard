package utils

import "os"

type FileUtils struct{}

func (f *FileUtils) Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !os.IsNotExist(err)
}
