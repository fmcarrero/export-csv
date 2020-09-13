package infrastructure_container

import "github.com/fmcarrero/export-csv/api/infrastructure/controller/movie"

func GetMovieController() *movie.ControllerMovie {
	return &movie.ControllerMovie{}
}
