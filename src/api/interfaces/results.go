package interfaces

type Results struct {
	Popularity  float32 `json:"popularity"`
	VoteCount   int16   `json:"vote_count"`
	Video       bool    `json:"video"`
	PosterPath  string  `json:"poster_path"`
	Title       string  `json:"title"`
	VoteAverage float32 `json:"vote_average"`
	Id          int32   `json:"id"`
	Overview    string  `json:"overview"`
}
