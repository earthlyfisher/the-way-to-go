package hellolib

import "testing"

func TestMax(t *testing.T) {
	const x, y = 2, 3
	max := Max(x, y)
	if max == x {
		t.Error("Max错误！")
	}
}
