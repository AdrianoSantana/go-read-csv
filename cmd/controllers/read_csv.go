package controller

import (
	"encoding/csv"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/AdrianoSantana/go-read-csv/cmd/models"
	"github.com/AdrianoSantana/go-read-csv/cmd/repositories"
)

var fileNotFound error = errors.New("Não foi possivel encontrar um arquivo com o caminho especificado")

func CreateData(fp string, r repositories.MovieRepository) (int, error) {
	movies, err := ReadCSV(fp)
	if err != nil {
		return 0, err
	}

	insertedMovies := r.Insert(movies)
	if err != nil {
		return 0, err
	}
	
	return insertedMovies, nil
}

func ReadCSV(fp string) (movies []models.Movie, err error) {
	_, err = os.Stat(fp)
	if err != nil {
		err = fileNotFound
		return
	}

	csvFile, err := os.Open(fp)
	if err != nil {
		return
	}
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return 
	}

	movies = createListOfMovies(lines)
	return
}

func createListOfMovies(lines [][]string) []models.Movie {
	movies := []models.Movie{}
	for _, line := range lines {
		rawId := line[0]
		title := line[1]
		genre := line[2]

		id, err := strconv.Atoi(rawId)
		if err != nil {
			continue
		}

		year, err := getYearOfMovie(title)
		if err != nil {
			continue
		}

		mv := createMovie(id, title, genre, year)
		movies = append(movies, mv)
	}
	return movies
}

func createMovie(id int, title, genre string, year int) models.Movie {
	return models.Movie{ID: id, Title: title, Genres: genre, Year: year }
}

func getYearOfMovie(title string) (year int, err error) {
	r, err := regexp.Compile(`\((.*?)\)`)
	if err != nil {
		return 0, err
	}
	matchAll := r.FindAllString(title, -1)
	for _, element := range matchAll {
		element = strings.Trim(element, "(")
		element = strings.Trim(element, ")")
		year, err = strconv.Atoi(element)
		if err != nil {
			return
		}
	}
	return
}
