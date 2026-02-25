package utils

import (
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	sourceFile, _ := os.Open(src)
	defer sourceFile.Close()
	destFile, _ := os.Create(dst)
	defer destFile.Close()
	_, err := io.Copy(destFile, sourceFile)
	return err
}
