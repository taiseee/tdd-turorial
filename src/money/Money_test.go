package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPlusReturnSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum := result.(Sum)
	assert.Equal(t, sum.augent, five)
	assert.Equal(t, sum.addend, five)
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), result)
}