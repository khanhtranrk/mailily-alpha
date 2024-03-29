package account_repository

import (
	"database/sql"
  "github.com/khanhtranrk/mailily-alpha/struct/mail"
)

func getAll(db *sql.DB) ([]mail.Account, error) {
  var accounts []mail.Account
  query := `SELECT id, host, port, username, password FROM accounts`

  rows, err := db.Query(query)
  if err != nil {
    return accounts, err
  }

  for rows.Next() {
    var account mail.Account
    if err := rows.Scan(&account.host, &account.port, &account.username, &account.password); err != nil {
      return accounts, err
    }
    accounts = append(accounts, account)
  }

  return accounts, nil
}
