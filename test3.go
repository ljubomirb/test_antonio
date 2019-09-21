package main

import (
	"os"
	"path/filepath"
)

//FindFirstLongestFileNameInDir returns the FIRST longest filename within folder (including subfolders)
func FindFirstLongestFileNameInDir(directoryPath string) (string, error) {
	var currentMaxFilenameLen int
	var currentMaxFilename string

	slashedPath := filepath.ToSlash(directoryPath)

	err := filepath.Walk(slashedPath, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			filename := info.Name()
			filenameLen := len(filename)

			if filenameLen > currentMaxFilenameLen {
				currentMaxFilename = filename
				currentMaxFilenameLen = filenameLen
			}
		}

		return nil
	})

	return currentMaxFilename, err
}
