package money

type Sum struct {
	augent Expression
	addend Expression
}

func (s Sum) Augent() Expression {
	return s.augent
}

func (s Sum) Addend() Expression {
	return s.addend
}

func (s Sum) Reduce(bank Bank, to string) Money {
	amount := s.augent.Reduce(bank, to).Amount() + s.addend.Reduce(bank, to).Amount()
	return NewMoney(amount, to)
}

func (s Sum) Plus(addend Expression) Expression {
	return NewSum(s, addend)
}

func (s Sum) Times(multiplier int) Expression {
	return NewSum(s.augent.Times(multiplier), s.addend.Times(multiplier))
}

func NewSum(augent Expression, addend Expression) Sum {
	return Sum{augent: augent, addend: addend}
}
