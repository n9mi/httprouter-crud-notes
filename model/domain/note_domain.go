package domain

type Note struct {
	Id       int
	Title    string
	Body     string
	Category string
}

type NoteWithCategoryId struct {
	Id         int
	Title      string
	Body       string
	IdCategory int
}
