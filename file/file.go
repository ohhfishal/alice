package file

import (
	"io"
	"os"
)

type Saveable interface {
	Save(writer io.Writer) error
}

func toFile(saveable Saveable, path string, flag int, perm os.FileMode) error {
	file, err := os.OpenFile(path, flag, perm)
	defer file.Close()
	if err != nil {
		return err
	}
	return saveable.Save(file)
}

func AppendToFile(saveable Saveable, path string) error {
	return toFile(saveable, path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
}

func ToFile(saveable Saveable, path string) error {
	return toFile(saveable, path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}
