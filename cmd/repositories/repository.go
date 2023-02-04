package repositories

import "github.com/AdrianoSantana/go-read-csv/cmd/models"

type MovieRepository interface {
	Insert(movies []models.Movie) int
}
