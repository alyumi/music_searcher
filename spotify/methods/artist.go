package methods

type Artist struct {
	ID string `json:"id"`

	Name   string   `json:"name"`
	Genres []string `json:"genres"`
}
