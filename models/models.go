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

type Sum struct {
	Val_1 int
	Val_2 int
}

type PodMetadata struct {
	Node              string `json:"node"`
	PodName           string `json:"podName"`
	PodIP             string `json:"podIp"`
	PodNamespace      string `json:"podNamespace"`
	PodServiceAccount string `json:"podServiceAccount"`
}
