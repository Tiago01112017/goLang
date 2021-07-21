package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID             string
	Name           string
	Number         string
	ExpirationMoth int32
	ExpirationYear int32
	CVV            int32
	Balance        float64 //Saldo
	Limit          float64
	CreatedAt      time.Time
}

func NewCredtCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()
	return c
}
