package dexpenses

type PaymentMethod int

const (
	Debit   PaymentMethod = 0
	Credit
	Cash
	Unknown
)
