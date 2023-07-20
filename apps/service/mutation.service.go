package service

import (
	"database/sql"
	"github.com/MandoPurba/rest-api/config"
)

type MutationDTO struct {
	FromAccount string
	ToAccount   string
	Amount      int64
	Balance     int64
}

func Transfer(dto MutationDTO) (bool, error) {
	tx, err := config.DB.Begin()
	if err != nil {
		return false, err
	}
	// GET ID ACCOUNT
	var (
		from int64
		to   int64
	)
	err = tx.QueryRow(`
		SELECT
			(SELECT id FROM accounts WHERE account_number = $1) AS from_account,
			(SELECT id FROM accounts WHERE account_number = $2) AS to_account;
	`, dto.FromAccount, dto.ToAccount).Scan(&from, &to)

	// MAKE TRANSFERS
	_, err = tx.Exec(`
		UPDATE amounts a SET balance = a.balance - $1 WHERE account_id = $2
	`, dto.Amount, from)

	_, err = tx.Exec(`
		UPDATE amounts a SET balance = a.balance + $1 WHERE account_id = $2
	`, dto.Balance, to)

	// COMMIT
	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func Amount(accountNumber string) (int64, int64, error) {
	var (
		balance int64
		tax     int64
	)
	row := config.DB.QueryRow(`
		SELECT am.balance, at.tax FROM amounts am
		JOIN accounts ac ON ac.id = am.account_id
		JOIN account_types at ON ac.account_type = at.id
		WHERE ac.account_number = $1
	`, accountNumber)
	err := row.Scan(&balance, &tax)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, err
		}
	}
	return balance, tax, nil
}
