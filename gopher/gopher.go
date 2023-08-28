package gopher

type Gopher struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Hired      bool   `json:"hired"`
	Profession string `json:"profession"`
}
