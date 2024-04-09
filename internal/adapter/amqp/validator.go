package amqp

import "strings"

type Validator struct {
}

func NewValidator() *Validator {
  return &Validator{}
}

func (v *Validator) Validate(mes string) (string, error) {
  if mes[preIndex:index] == "account" {
    preIndex := index + 1
    index := strings.Index(mes, " ")

    if mes[preIndex:index] == "list" {
      return "", nil
    }
  }
}
