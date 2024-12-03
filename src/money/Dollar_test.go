package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)

	assert.Equal(t, NewDollar(10), five.Times(2))

	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	five := NewDollar(5)

	// 実装方法がわかっている場合はこれだけでもいい
	assert.Equal(t, NewDollar(15), five.Times(3))

	// 設計アイデアが浮かばない場合は二つ目の実例を追加し三角測量を行ってみる
	// それによりアイデアが浮かぶ可能性がある
	assert.NotEqual(t, NewDollar(6), five.Times(3))
}
