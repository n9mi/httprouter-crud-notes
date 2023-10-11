package web

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,min=2,max=200"`
}
