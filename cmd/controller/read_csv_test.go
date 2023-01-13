package controller

import (
	"reflect"
	"testing"

	"github.com/AdrianoSantana/go-read-csv/cmd/models"
)

func TestReadCsv(t *testing.T) {
	t.Run("Should return an error if no file is found it", func(t *testing.T) {
		filePath := "movies"
		_, err := ReadCSV(filePath)

		expected := fileNotFound
		if err != expected {
			t.Errorf("Result: '%s' Expected: '%s'", err, expected)
		}
	})

	t.Run("Should return an array of movies from csv", func(t *testing.T) {
		filepath := "../../test"
		result, err := ReadCSV(filepath)

		if err != nil {
			t.Fatal("Not expected an error!")
		}

		expected := []models.Movie{
			{ID: 1, Title: "Kein Bund für's Leben (2007)", Genres: "Comedy"},
			{ID: 2, Title: "Feuer, Eis & Dosenbier (2002)", Genres: "Comedy"},
			{ID: 3, Title: "The Pirates (2014)", Genres: "Adventure"},
			{ID: 4, Title: "Rentun Ruusu (2001)", Genres: "(no genres listed)"},
			{ID: 5, Title: "Innocence (2014)", Genres: "Adventure|Fantasy|Horror"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result: '%v' Expected: '%v'", result, expected)
		}
	})
}
