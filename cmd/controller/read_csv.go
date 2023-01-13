package controller

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/AdrianoSantana/go-read-csv/cmd/models"
)

var fileNotFound error = errors.New("Não foi possivel encontrar um arquivo com o caminho especificado")

func ReadCSV(filePath string) ([]models.Movie, error) {
	movies := []models.Movie{}

	_, err := os.Stat(filePath)
	if err != nil {
		return nil, fileNotFound
	}

	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			continue
		}
		mv := models.Movie{
			ID:     id,
			Title:  line[1],
			Genres: line[2],
		}
		movies = append(movies, mv)
	}
	return movies, nil
}
