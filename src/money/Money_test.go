package money

import (
	"testing"
)

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	if !sum.Equals(reduced) {
		t.Errorf("Error")
	}
}
