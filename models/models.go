package models

import (
	"time"

	"github.com/jmoiron/sqlx"
	// postgres db driver
	_ "github.com/lib/pq"
)

// RawBankStatement is a struct representing a
// db object of the same name
type RawBankStatement struct {
	ID              int       `json:"id" db:"id"`
	Posted          time.Time `json:"posted" db:"posted"`
	TransactionDesc string    `json:"transaction_desc" db:"transaction_desc"`
	Debit           float64   `json:"debit" db:"debit"`
	Credit          float64   `json:"credit" db:"credit"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	Provider        string    `json:"provider" db:"provider"`
}

// AllRawBankStatements pull all rawbankstatements from the db
func AllRawBankStatements(db *sqlx.DB) ([]RawBankStatement, error) {
	rows, err := db.Queryx("SELECT * FROM raw_bank_statements")
	if err != nil {
		return nil, err
	}
	var rawStatements []RawBankStatement
	defer rows.Close()

	for rows.Next() {
		var rawStatement RawBankStatement
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
