package entity

type Account struct {
	Id            int64       `json:"id,omitempty"`
	AccountNumber string      `json:"account_number,omitempty"`
	CreatedAt     string      `json:"created_at,omitempty"`
	Type          AccountType `json:"type,omitempty"`
}

type AccountType struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Tax         int32  `json:"tax,omitempty"`
}
