package repositories

import (
	"testing"

	"github.com/AdrianoSantana/go-read-csv/cmd/models"
)

func Test_Sqlite_Repository(t *testing.T) {
	t.Run("Should insert movies", func(t *testing.T) {
		input := []models.Movie{
			{ID: 1, Title: "Kein Bund für's Leben (2007)", Genres: "Comedy", Year: 2007},
			{ID: 2, Title: "Feuer, Eis & Dosenbier (2002)", Genres: "Comedy", Year: 2002},
			{ID: 3, Title: "The Pirates (2014)", Genres: "Adventure", Year: 2014},
			{ID: 4, Title: "Rentun Ruusu (2001)", Genres: "(no genres listed)", Year: 2001},
			{ID: 5, Title: "Innocence (2014)", Genres: "Adventure|Fantasy|Horror", Year: 2014},
		}

		sqliteRepository := MovieRepositorySqlite{
			DatabaseConn: "./test-database.db",
		}
		insertedRows := sqliteRepository.Insert(input)

		expected := len(input)

		if insertedRows != expected {
			t.Errorf("Result: %d, Expected: %d", expected, insertedRows)
		}
	})
}
