package dexpenses

import (
	"github.com/rhymond/go-money"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"time"
)

type PersistentMoney struct {
	Amount   float64
	Currency string
}

func AsPersistentMoney(m *money.Money) PersistentMoney {
	return PersistentMoney{Amount: float64(m.Amount()) / 100, Currency: m.Currency().Code}
}

type Receipt struct {
	ID            objectid.ObjectID `bson:"_id,omitempty"`
	Date          time.Time
	Time          time.Time
	Total         PersistentMoney
	PaymentMethod PaymentMethod     `bson:"paymentMethod"`
	Category      string
}
