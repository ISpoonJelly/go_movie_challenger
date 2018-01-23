package challenges

type Challenge struct {
  Page          int    `json:"page"`
  Total_results int       `json:"total_results"`
  Total_pages   int       `json:"total_pages"`
  Results       []Movie   `json:"results"`
}

type Movie struct {
  Vote_count          int       `json:"vote_count"`
  Title               string    `json:"title"`
  Genre_ids           []int     `json:"genre_ids"`
  ID                  int       `json:"id"`
  Video               bool      `json:"video"`
  Vote_average        float32   `json:"vote_average"`
  Popularity          float32   `json:"popularity"`
  Poster_path         string    `json:"poster_path"`
  Original_language   string    `json:"original_language"`
  Original_title      string    `json:"original_title"`
  Backdrop_path       string    `json:"backdrop_path"`
  Adult               bool      `json:"adult"`
  overview            string    `json:"release_date"`
  release_date        string    `json:"release_date"`
}

type Genre_ids struct {
  Genre_ids   []int   `json:"genre_ids"`
}
