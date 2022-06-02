package domain

import (
	"regexp"

	"github.com/marcosvieirajr/payment/internal/app"
)

type Account struct {
	ID             int64
	DocumentNumber string
}

func (a *Account) Validate() error {
	match, _ := regexp.MatchString("^[0-9]{11}$", a.DocumentNumber)
	if !match {
		return app.ErrDocumentNumberIsInvalid
	}

	return nil
}
