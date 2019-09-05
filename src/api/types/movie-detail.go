package types

type MovieDetail struct {
	Homepage         string  `json:"homepage"`
	Id               int32   `json:"id"`
	ImdbId           string  `json:"imdb_id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
}
