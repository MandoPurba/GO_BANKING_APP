package service

import (
	"database/sql"
	"github.com/MandoPurba/rest-api/apps/entity"
	"github.com/MandoPurba/rest-api/config"
)

type UserDTO struct {
	Email     string
	Password  string
	Hash      string
	FirstName string
	LastName  string
	Address   string
}

func GetALlUsers() ([]entity.User, error) {
	var users []entity.User
	rows, err := config.DB.Query(`
		SELECT 
				u.id, u.email, u.created_at,
				a.id, a.account_number, a.created_at,
				t.id, t.description, t.name, t.tax
		FROM users u
		JOIN accounts a on a.id = u.account_id
		JOIN account_types t on t.id = a.account_type
		WHERE active = true
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Email, &user.CreatedAt, &user.Account.Id, &user.Account.AccountNumber, &user.Account.CreatedAt, &user.Account.Type.Id, &user.Account.Type.Name, &user.Account.Type.Description, &user.Account.Type.Tax)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int64) (*entity.User, error) {
	var user entity.User
	row := config.DB.QueryRow(`
		SELECT 
				u.id, u.email, u.created_at,
				a.id, a.account_number, a.created_at,
				t.id, t.description, t.name, t.tax
		FROM users u
		JOIN accounts a on a.id = u.account_id
		JOIN account_types t on t.id = a.account_type
		WHERE active = true
		AND u.id = $1
	`, id)
	err := row.Scan(&user.Id, &user.Email, &user.CreatedAt, &user.Account.Id, &user.Account.AccountNumber, &user.Account.CreatedAt, &user.Account.Type.Id, &user.Account.Type.Name, &user.Account.Type.Description, &user.Account.Type.Tax)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user UserDTO) (bool, error) {
	tx, err := config.DB.Begin()
	if err != nil {
		return false, err
	}
	// INSERT USER
	var userId int64
	err = tx.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id", user.Email, user.Hash).Scan(&userId)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	//INSERT DETAIL USERS
	_, err = tx.Exec("INSERT INTO detail_users (first_name, last_name, user_id, address) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, userId, user.Address)
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
