package service

import (
	"github.com/MandoPurba/rest-api/config"
)

type AccountDTO struct {
	AccountType int
	Currency    int
	UserId      int
}

func CreateAccount(accountNumber string, accountType, currency, userId int) (bool, error) {
	tx, err := config.DB.Begin()
	if err != nil {
		return false, err
	}

	var accountId int64
	err = tx.QueryRow("INSERT INTO  accounts (account_number, account_type) VALUES ($1, $2) RETURNING id", accountNumber, accountType).Scan(&accountId)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = tx.Exec("INSERT INTO amounts (account_id, balance, currency_code) VALUES ($1,0,$2);", accountId, currency)
	if err != nil {
		return false, err
	}

	_, err = tx.Exec("UPDATE users SET account_id = $1, active = true WHERE id = $2", accountId, userId)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}
