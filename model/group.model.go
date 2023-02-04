package model

type GroupAttr struct {
	Phone   []string `json:"phone,omitempty"`
	Company []string `json:"company"`
}
type Group struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Attributes GroupAttr `json:"attributes"`
}
