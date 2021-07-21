package repository

import (
	"database/sql"

	"github.com/Tiago-Alves-Dev/codebank/domain"
)

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{db: db}
}

func (t *TransactionRepositoryDB) SaveTransaction(transaction domain.Transaction, credirtCard domain.CreditCard) error {
	stmt, err := t.db.Prepare(`
	insert into transations( id, credit_card_id, amount, status, description, store, creatd_at )
	values( $1, $2, $3, $4, $5, $6, $7 )
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardId,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	)

	if err != nil {
		return err
	}

	if transaction.Status == "Approved" {
		err = t.updateBalance(credirtCard)
		if err != nil {
			return err
		}
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil

}

func (t *TransactionRepositoryDB) updateBalance(credirtCard domain.CreditCard) error {
	_, err := t.db.Exec("update credit_cards set balance = $1 where id = $2",
		credirtCard.Balance, credirtCard.ID)

	if err != nil {
		return err
	}
	return nil
}
