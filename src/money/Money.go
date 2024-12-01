package money

// type Accessor interface {
// 	Amount() int
// 	Currency() string
// }

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *Money {
	return &Money{amount: amount, currency: currency}
}

// NewDollar is constructor of Dollar.
func NewDollar(a int) *Money {
	return NewMoney(a, "USD")
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) *Money {
	return NewMoney(a, "CHF")
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.amount*multiplier, m.currency)
}

func (m *Money) Plus(addend *Money) *Money {
	return NewMoney(m.amount+addend.amount, m.currency)
}

func (m *Money) Equals(a *Money) bool {
	return m.amount == a.amount && m.currency == a.currency
}

// func (m *Money) Amount() int {
// 	return m.amount
// }

// func (m *Money) Currency() string {
// 	return m.currency
// }

func (m *Money) Dollar() *Money {
	return NewDollar(m.amount)
}

func (m *Money) Franc() *Money {
	return NewFranc(m.amount)
}
