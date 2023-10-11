package web

type NoteResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category string `json:"category"`
}

type NoteShortResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
}

type NoteRequest struct {
	Id         int    `json:"id"`
	Title      string `json:"title" validate:"required,min=2,max=100"`
	Body       string `json:"body" validate:"required,min=2,max=255"`
	IdCategory int    `json:"id_category" validate:"required,min=1"`
}
