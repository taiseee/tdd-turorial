package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)

	if !five.Times(2).Equals(NewDollar(10)) {
		t.Errorf("Error")
	}

	if !five.Times(3).Equals(NewDollar(15)) {
		t.Errorf("Error")
	}
}

func TestEquality(t *testing.T) {
	five := NewDollar(5)

	// 実装方法がわかっている場合はこれだけでもいい
	product := five.Times(3)
	want := NewDollar(15)
	if !product.Equals(want) {
		t.Errorf("Product.amount = %d, want %d", product.amount, want.amount)
	}

	// 設計アイデアが浮かばない場合は二つ目の実例を追加し三角測量を行ってみる
	// それによりアイデアが浮かぶ可能性がある
	six := NewDollar(6)
	if product.Equals(six) {
		t.Errorf("Product.amount = %d, but not want %d", product.amount, six.amount)
	}
}
