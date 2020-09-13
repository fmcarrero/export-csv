package port

import "github.com/fmcarrero/export-csv/api/domain"

type MovieRepository interface {
	GetAllMovies() ([]domain.Movie, error)
}
