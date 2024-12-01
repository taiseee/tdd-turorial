package money

import (
	"testing"
)

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)

	if !five.Times(2).Equals(NewFranc(10)) {
		t.Errorf("Error")
	}

	if !five.Times(3).Equals(NewFranc(15)) {
		t.Errorf("Error")
	}
}
