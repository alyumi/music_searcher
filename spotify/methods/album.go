package methods

type Album struct {
	ID string `json:"id"`

	Name        string   `json:"name"`
	ReleaseDate string   `json:"release_date"`
	Genres      []string `json:"genres"`
}
