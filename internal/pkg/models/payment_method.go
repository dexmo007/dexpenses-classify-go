package dexpenses

//PaymentMethod enum
type PaymentMethod int

const (
	//Debit payment method debit / EC cash
	Debit PaymentMethod = 0
	//Credit payment method credit card
	Credit
	//Cash payment method by cash
	Cash
	//Unknown unknown payment method
	Unknown
)
