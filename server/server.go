package server

type (
	Category struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	Categories []Category
)

