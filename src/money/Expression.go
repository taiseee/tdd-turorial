package money

type Expression interface {
	Reduce(Bank, string) Money
	Plus(Expression) Expression
	Times(float64) Expression
}
