package controller

import (
	"reflect"
	"testing"

	"github.com/AdrianoSantana/go-read-csv/cmd/models"
)

func Test_Read_csv(t *testing.T) {
	t.Run("Should read a csv", func(t *testing.T) {
		expected := []models.Movie{
			{ID: 1, Title: "Kein Bund für's Leben (2007)", Genres: "Comedy", Year: 2007},
			{ID: 2, Title: "Feuer, Eis & Dosenbier (2002)", Genres: "Comedy", Year: 2002},
			{ID: 3, Title: "The Pirates (2014)", Genres: "Adventure", Year: 2014},
			{ID: 4, Title: "Rentun Ruusu (2001)", Genres: "(no genres listed)", Year: 2001},
			{ID: 5, Title: "Innocence (2014)", Genres: "Adventure|Fantasy|Horror", Year: 2014},
		}
		result, err := ReadCSV("../../test")
		if err != nil {
			t.Errorf("Not expected an error")
		}

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Result: %v, Expected: %v", result, expected)
		}

	})

	t.Run("Should identify the year of a movie by title", func(t *testing.T) {
		testTable := []struct {
			title    string
			expected int
		}{
			{title: "any_title (2023)", expected: 2023},
			{title: "other_title (1994)", expected: 1994},
			{title: "another_title", expected: 0},
		}

		for _, test := range testTable {
			result, err := getYearOfMovie(test.title)
			if err != nil {
				t.Errorf("Not expected an error")
			}

			if result != test.expected {
				t.Errorf("Result: %d, Expected: %d", result, test.expected)
			}
		}
	})
}
