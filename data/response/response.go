package response

type NoteResponse struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type BookResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
