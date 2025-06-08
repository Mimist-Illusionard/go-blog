package dto

type CommentResponse struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
