package models

type CatFact struct {
	ID        string `json:"_id"`
	V         int    `json:"__v"`
	Text      string `json:"text"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
	User      string `json:"user"`
	Deleted   bool   `json:"deleted"`
}
