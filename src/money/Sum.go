package money

type Sum struct {
	augent Money
	addend Money
}

func (s Sum) Augent() Money {
	return s.augent
}

func (s Sum) Addend() Money {
	return s.addend
}

func (s Sum) Reduce(to string) Money {
	amount := s.augent.Amount() + s.addend.Amount()
	return NewMoney(amount, to)
}

func NewSum(augent Money, addend Money) Sum {
	return Sum{augent: augent, addend: addend}
}
