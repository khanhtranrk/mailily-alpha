package amqp

import (
	"fmt"
	"strings"

	"github.com/khanhtranrk/mailily-alpha/internal/core/port"
)

type AccountHandle struct {
  svc port.AccountService
}

func NewAccountHandle(svc port.AccountService) *AccountHandle {
  return &AccountHandle{
    svc: svc,
  }
}

func (ah *AccountHandle) Exec(mes string) (string error) {
  index := strings.Index(mes, " ")
  if mes[0:index] == "list" {
    return ListAccounts(mes[0:index + 1], ah.svc)
  }

  return nil, fmt.Errorf("This command is wrong - Unknow '%s'", mes[0:index])
}

func ListAccounts(mes string, svc port.AccountService) (string, error) {
  index := strings.Index(mes, " ")
  if mes[0:index] != 
}
