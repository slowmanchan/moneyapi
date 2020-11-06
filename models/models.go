package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type RawBankStatment struct {
	ID              int       `json:"id" db:"id"`
	Posted          time.Time `json:"posted" db:"posted"`
	TransactionDesc string    `json:"transaction_desc" db:"transaction_desc"`
	Debit           float64   `json:"debit" db:"debit"`
	Credit          float64   `json:"credit" db:"credit"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	Provider        string    `json:"provider" db:"provider"`
}

func AllRawBankStatements(db *sqlx.DB) ([]RawBankStatment, error) {
	rows, err := db.Queryx("SELECT * FROM raw_bank_statements")
	if err != nil {
		return nil, err
	}
	var rawStatements []RawBankStatment
	defer rows.Close()

	for rows.Next() {
		var rawStatement RawBankStatment
		if err := rows.StructScan(&rawStatement); err != nil {
			return nil, err
		}
		rawStatements = append(rawStatements, rawStatement)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rawStatements, nil
}
