package money

type Pair struct {
	from string
	to   string
}

func NewPair(from string, to string) Pair {
	return Pair{from: from, to: to}
}

type Rates map[Pair]int

type Bank struct {
	rates Rates
}

func (b Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(b, to)
}

func (b *Bank) AddRate(from string, to string, rate int) {
	b.rates[NewPair(from, to)] = rate
}

func (b *Bank) Rate(from string, to string) int {
	if from == to {
		return 1
	}
	return b.rates[NewPair(from, to)]
}

func NewBank() *Bank {
	return &Bank{rates: make(Rates)}
}
