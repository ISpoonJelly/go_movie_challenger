package movies

type APIParams struct {
	APIKey string `url:"api_key"`
}

type Genres struct {
	Result []Genre `json:"genres"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Discover struct {
	Page         int             `json:"page"`
	TotalResults int             `json:"total_results"`
	TotalPages   int             `json:"total_pages"`
	Movie        []DiscoverMovie `json:"results"`
}

type DiscoverMovie struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Poster      string `json:"poster_path,omitempty"`
	Rating      int    `json:"vote_average,omitempty"`
	Popularity  int    `json:"popularity,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	Overview    string `json:"overview,omitempty"`
}

type TmdbError struct {
	Message string `json:"status_message"`
	Code    int    `json:"status_code"`
}
