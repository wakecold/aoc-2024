package main

import "testing"

func TestCheckVals(t *testing.T) {
	tc := []int{97, 75, 47, 29, 13}
	rulesBack := []int{13, 61, 47, 29, 53, 75}
	rulesFront := []int{}

	if !checkVals(tc, rulesBack, rulesFront, 0) {
		t.Fail()
	}
}
