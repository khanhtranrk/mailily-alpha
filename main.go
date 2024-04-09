package main

import (
	"log"

	"github.com/khanhtranrk/mailily-alpha/functions/sender"
)

func main() {
  err := sender.Send()
  if err != nil {
    log.Fatal(err)
  }
}
