package db

import (
	"database/sql"

	"github.com/Htet-Shine-Htwe/bank_go/app/types"
)

type Storage interface {
	createAccount(*types.Account) error
	getAccount(int) error
	updateAccount(*types.Account) error
	GetAccountById(int) (*types.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connectStr := "user=postgres dbname=go_bank password=gobank sslmode=verify-full"

	db, err := sql.Open("postgres", connectStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}
