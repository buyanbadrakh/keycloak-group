package model

type UserAttr struct {
	Phone   []string `json:"phone,omitempty"`
	Company []string `json:"company"`
}
type User struct {
	ID               string   `json:"id"`
	FirstName        string   `json:"firstName"`
	LastName         string   `json:"lastName"`
	UserName         string   `json:"username"`
	Email            string   `json:"email"`
	Attributes       UserAttr `json:"attributes"`
	Enabled          bool     `json:"enabled"`
	EmailVerified    bool     `json:"emailVerified"`
	CreatedTimestamp int64    `json:"createdTimestamp"`
}
