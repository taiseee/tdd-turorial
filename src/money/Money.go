package money

type Accessor interface {
	Amount() int
	Name() string
}

type Money struct {
	amount int
	name   string
}

func NewMoney(amount int, name string) *Money {
	return &Money{amount: amount, name: name}
}

// NewDollar is constructor of Dollar.
func NewDollar(a int) *Money {
	return NewMoney(a, "Dollar")
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) *Money {
	return NewMoney(a, "Franc")
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (m *Money) Times(multiplier int) *Money {
	return &Money{
		amount: m.amount * multiplier,
		name:   m.name,
	}
}

func (m *Money) Equals(a Accessor) bool {
	return m.Amount() == a.Amount() && m.Name() == a.Name()
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Name() string {
	return m.name
}

func (m *Money) Dollar() *Money {
	return NewDollar(m.amount)
}

func (m *Money) Franc() *Money {
	return NewFranc(m.amount)
}
