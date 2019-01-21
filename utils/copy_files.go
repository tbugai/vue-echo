package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Copy(src, dest string) error {
	info, err := os.Lstat(src)
	if err != nil {
		return err
	}

	return copy(src, dest, info)
}

func copy(src, dest string, info os.FileInfo) error {
	if info.IsDir() {
		return dcopy(src, dest, info)
	}

	return fcopy(src, dest, info)
}

func fcopy(src, dest string, info os.FileInfo) error {
	if err := os.MkdirAll(filepath.Dir(dest), os.ModePerm); err != nil {
		return err
	}

	directory, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer directory.Close()

	if err = os.Chmod(directory.Name(), info.Mode()); err != nil {
		return err
	}

	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(directory, file)

	return err
}

func dcopy(src, dest string, info os.FileInfo) error {
	if err := os.MkdirAll(dest, info.Mode()); err != nil {
		return err
	}

	contents, err := ioutil.ReadDir(src)
	if err != nil {
		return nil
	}

	for _, content := range contents {
		sourceFile, destinationFile := filepath.Join(src, content.Name()), filepath.Join(dest, content.Name())
		if err := copy(sourceFile, destinationFile, content); err != nil {
			return err
		}
	}

	return nil
}
