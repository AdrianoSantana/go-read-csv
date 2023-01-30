package repositories

import (
	"database/sql"
	"log"

	"github.com/AdrianoSantana/go-read-csv/cmd/models"

	_ "github.com/mattn/go-sqlite3"
)

type MovieRepositorySqlite struct{}

func (mv MovieRepositorySqlite) Insert(movies []models.Movie) (insertedMovies int) {
	db, err := sql.Open("sqlite3", "./database/movies-csv.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = isConnected(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies {
		smt, err := db.Prepare("insert into movies(title, genres) VALUES(?, ?)")
		if err != nil {
			log.Fatal(err)
			continue
		}
		_, err = smt.Exec(movie.Title, movie.Genres)
		if err != nil {
			log.Fatal(err)
			continue
		}
		insertedMovies += 1
	}
	return
}

func createOrGetTable(db *sql.DB) (error) {
	_, err := db.Query("CREATE TABLE IF NOT EXISTS movies (id integer PRIMARY KEY, title TEXT NOT NULL, genres TEXT NOT NULL)")
	if err != nil {
		return err
	}
	return nil
}

func isConnected(db *sql.DB) (error) {
	return db.Ping()
}