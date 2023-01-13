package controller

import (
	"os"
)

func ReadCSV(filePath string) (string, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	return "", nil
}
