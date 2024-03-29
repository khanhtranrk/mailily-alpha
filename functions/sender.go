package sender

import (
  "fmt"
  "github.com/khanhtranrk/mailily-alpha/repositories/account_repository"
)

type Message struct {
    From string
    To string
    Subject string
    Body string
}

func Send(message Message) error {
  fmt.Println("Sent!")

  return nil
}
