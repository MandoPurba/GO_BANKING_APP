package entity

type User struct {
	Id        int     `json:"id,omitempty"`
	Email     string  `json:"email,omitempty"`
	Password  string  `json:"-"`
	CreatedAt string  `json:"created_at,omitempty"`
	Account   Account `json:"account,omitempty"`
}

type DetailUser struct {
	Id        int64  `json:"id,omitempty"`
	User      User   `json:"user,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Address   string `json:"address,omitempty"`
}
