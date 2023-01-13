package controller

import (
	"errors"
	"os"
)

var fileNotFound error = errors.New("Não foi possivel encontrar um arquivo com o caminho especificado")

func ReadCSV(filePath string) (string, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return "", fileNotFound
	}
	return "", nil
}
