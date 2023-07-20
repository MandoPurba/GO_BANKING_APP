package entity

type Amount struct {
	Id       int64    `json:"id,omitempty"`
	Account  Account  `json:"account,omitempty"`
	Balance  int64    `json:"balance,omitempty"`
	Currency Currency `json:"currency,omitempty"`
}

type Currency struct {
	Id      int64  `json:"id,omitempty"`
	Code    string `json:"code,omitempty"`
	Country string `json:"country,omitempty"`
}
