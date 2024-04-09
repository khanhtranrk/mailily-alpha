package port

import "github.com/khanhtranrk/mailily-alpha/internal/core/domain"

type AccountRepository interface {
  CreateAccount() ([]domain.Account, error)
}

type AccountService interface {
  ListAccounts() ([]domain.Account, error)
}
