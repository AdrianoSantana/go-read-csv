package repositories

import (
	"database/sql"
	"log"
	"sync"

	"github.com/AdrianoSantana/go-read-csv/cmd/models"

	_ "github.com/mattn/go-sqlite3"
)

var wg sync.WaitGroup
var mut sync.Mutex

type MovieRepositorySqlite struct {
	DatabaseConn string
}

func (mv MovieRepositorySqlite) Insert(movies []models.Movie) (insertedMovies int) {
	db, err := sql.Open("sqlite3", mv.DatabaseConn)
	if err != nil {
		log.Fatal(err)
	}
	err = createOrGetTable(db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = isConnected(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies {
		go saveMovieInDB(db, movie, &insertedMovies)
		wg.Add(1)
	}
	wg.Wait()
	return
}

func saveMovieInDB(db *sql.DB, movie models.Movie, insertedMovies *int) {
	defer wg.Done()
	mut.Lock()
	defer mut.Unlock()
	insertInDatabase(db, movie)
	*insertedMovies += 1
}

func insertInDatabase(db *sql.DB, movie models.Movie) {
	smt, err := db.Prepare("insert into movies(title, genres) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = smt.Exec(movie.Title, movie.Genres)
	if err != nil {
		log.Fatal(err)
	}
}

func createOrGetTable(db *sql.DB) error {
	_, err := db.Query("CREATE TABLE IF NOT EXISTS movies (id integer PRIMARY KEY, title TEXT NOT NULL, genres TEXT NOT NULL)")
	if err != nil {
		return err
	}
	return nil
}

func isConnected(db *sql.DB) error {
	return db.Ping()
}
