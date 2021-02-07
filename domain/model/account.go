package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"ownerName" valid:"notnull"`
	Bank      *Bank     `json:"bank" valid:"notnull"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (account *Account) validate() error {
	if _, err := govalidator.ValidateStruct(account); err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		Bank:      bank,
		OwnerName: ownerName,
		Number:    number,
	}

	account.Id = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	if err := account.validate(); err != nil {
		return nil, err
	}

	return &account, nil
}
