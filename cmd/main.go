package main

import (
	"fmt"
	"log"

	controller "github.com/AdrianoSantana/go-read-csv/cmd/controllers"
	"github.com/AdrianoSantana/go-read-csv/cmd/repositories"
)

func main() {
	pathFile := "../movies"
	rp := repositories.MovieRepositorySqlite{
		DatabaseConn: "./database/movies-csv.db",
	}

	insertedMovies, err := controller.CreateData(pathFile, rp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %d movies\n", insertedMovies)
}
