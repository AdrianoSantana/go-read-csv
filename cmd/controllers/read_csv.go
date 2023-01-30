package controller

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/AdrianoSantana/go-read-csv/cmd/models"
	"github.com/AdrianoSantana/go-read-csv/cmd/repositories"
)

var fileNotFound error = errors.New("Não foi possivel encontrar um arquivo com o caminho especificado")

func ReadCSV(fp string, r repositories.MovieRepository) (int, error) {
	movies := []models.Movie{}

	_, err := os.Stat(fp)
	if err != nil {
		return 0, fileNotFound
	}

	csvFile, err := os.Open(fp)
	if err != nil {
		return 0, err
	}
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return 0, err
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

	insertedMovies := r.Insert(movies)
	if err != nil {
		return 0, err
	}

	return insertedMovies, nil
}
