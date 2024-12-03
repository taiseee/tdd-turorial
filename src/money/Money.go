package money

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) Money {
	return Money{amount: amount, currency: currency}
}

// NewDollar is constructor of Dollar.
func NewDollar(a float64) Money {
	return NewMoney(a, "USD")
}

// NewFranc is constructor of Dollar.
func NewFranc(a float64) Money {
	return NewMoney(a, "CHF")
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (m Money) Times(multiplier float64) Expression {
	return NewMoney(m.amount*multiplier, m.currency)
}

func (m Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount / rate, to)
}

func (m Money) Equals(a Money) bool {
	return m.amount == a.amount && m.currency == a.currency
}

func (m Money) Amount() float64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Dollar() Money {
	return NewDollar(m.amount)
}

func (m Money) Franc() Money {
	return NewFranc(m.amount)
}
