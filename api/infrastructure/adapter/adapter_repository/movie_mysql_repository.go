package adapter_repository

import (
	"database/sql"

	"github.com/fmcarrero/export-csv/api/domain"
)

type MovieMysqlRepository struct {
	DB *sql.DB
}

func (repository *MovieMysqlRepository) GetAllMovies() ([]domain.Movie, error) {
	query := "SELECT M.name, M.genre, M.assessment , M.synapses FROM movie M"
	rows, err := repository.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []domain.Movie
	for rows.Next() {
		var movie domain.Movie
		err = rows.Scan(&movie.Name, &movie.Genre, &movie.Assessment, &movie.Synapses)
		if err != nil {
			return movies, err
		}
		movies = append(movies, movie)
	}
	return movies, nil

}
