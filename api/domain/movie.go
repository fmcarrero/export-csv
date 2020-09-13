package domain

type Movie struct {
	Name       string
	Genre      string
	Assessment int64
	Synapses   string
}

func (movie *Movie) GetName() string {
	return movie.Name
}

func (movie *Movie) GetGenre() string {
	return movie.Genre
}
