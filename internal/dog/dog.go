package dog

type Dog struct {
	Colour string `json:"colour" db:"colour"`
	Breed  string `json:"breed" db:"breed"`
	Name   string `json:"name" db:"name"`
	Age    int    `json:"age" db:"age"`
}
