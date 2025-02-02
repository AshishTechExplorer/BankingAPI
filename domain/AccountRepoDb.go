package domain

import (
	"banking/logging"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepoDb struct {
	client *sqlx.DB
}

func (d AccountRepoDb) Save(a Account) (*Account, error) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logging.Info("Error while creating new account: " + err.Error())
		//return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logging.Info("Error while getting last insert id for new account: " + err.Error())
		//return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepoDb {
	return AccountRepoDb{dbClient}
}
