package main

import "testing"

func TestValidateAsc(t *testing.T) {
	tcs := []struct {
		name   string
		input  []int
		result bool
	}{
		{
			"first and second same",
			[]int{1, 1, 2, 3, 4, 5},
			true,
		},
		{
			"last two same",
			[]int{1, 2, 3, 4, 5, 5},
			true,
		},
		{
			"throw away first",
			[]int{5, 1, 2, 3, 4, 5},
			true,
		},
		{
			"throw away last",
			[]int{1, 2, 3, 4, 3},
			true,
		},
		{
			"test example fail",
			[]int{1, 2, 7, 8, 9},
			false,
		},
		{
			"happy path",
			[]int{1, 2, 3, 4, 5},
			true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.result != ValidateAsc(tc.input) {
				t.Fatalf("%s failed", tc.name)
			}
		})
	}
}

func TestValidateDesc(t *testing.T) {
	tcs := []struct {
		name   string
		input  []int
		result bool
	}{
		{
			"first and second same",
			[]int{5, 5, 4, 3, 2, 1},
			true,
		},
		{
			"last two same",
			[]int{5, 4, 3, 2, 1, 1},
			true,
		},
		{
			"throw away first",
			[]int{1, 5, 4, 3, 2, 1},
			true,
		},
		{
			"throw away last",
			[]int{5, 4, 3, 2, 1, 5},
			true,
		},
		{
			"test example fail",
			[]int{9, 7, 6, 2, 1},
			false,
		},
		{
			"happy path",
			[]int{9, 7, 6, 5},
			true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.result != ValidateDesc(tc.input) {
				t.Fatalf("%s failed", tc.name)
			}
		})
	}
}
