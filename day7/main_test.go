package main

import "testing"

func TestDPWithConcat(t *testing.T) {
	vals := []int{17, 8, 14}
	if !dpWithConcat(1, vals, 17, 192) {
		t.Fail()
	}
}
