package dto

type PostResponse struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
