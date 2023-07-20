package entity

type Mutation struct {
	Id           int64   `json:"id,omitempty"`
	FromAccount  Account `json:"from_account,omitempty"`
	ToAccount    Account `json:"to_account,omitempty"`
	Count        int64   `json:"price,omitempty"`
	MutationDate string  `json:"mutation_date,omitempty"`
}
