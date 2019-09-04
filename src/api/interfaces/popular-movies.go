package interfaces

type PopularMovies struct {
	Page         int       `json:"page"`
	TotalResults int16     `json:"total_results"`
	TotalPages   int       `json:"total_pages"`
	Results      []Results `json:"results"`
}
