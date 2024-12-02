package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()

		strs := strings.Split(line, " ")
		if len(strs) < 2 {
			panic("incorect len")
		}
		vals := make([]int, len(strs))
		for i, s := range strs {
			val, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			vals[i] = val
		}

		if ValidateDesc(vals) {
			result++
		} else {
			if ValidateAsc(vals) {
				result++
			}
		}

	}

	fmt.Println("Result: ", result)
}

func ValidateDesc(vals []int) bool {
	for i := range vals {
		if i == len(vals)-1 {
			break
		}

		diff := vals[i] - vals[i+1]
		if diff < 1 || diff > 3 {
			buf1 := make([]int, len(vals))
			copy(buf1, vals)
			arr1 := append(buf1[:i], buf1[i+1:]...)
			buf2 := make([]int, len(vals))
			copy(buf2, vals)
			arr2 := append(buf2[:i+1], buf2[i+2:]...)
			if validateDescNoIssues(arr1) || validateDescNoIssues(arr2) {
				return true
			}
		}
	}

	return validateDescNoIssues(vals)
}

func validateDescNoIssues(vals []int) bool {
	for i := range vals {
		if i == len(vals)-1 {
			break
		}
		diff := vals[i] - vals[i+1]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func ValidateAsc(vals []int) bool {
	for i := range vals {
		if i == len(vals)-1 {
			break
		}

		diff := vals[i+1] - vals[i]
		if diff < 1 || diff > 3 {
			buf1 := make([]int, len(vals))
			copy(buf1, vals)
			arr1 := append(buf1[:i], buf1[i+1:]...)
			buf2 := make([]int, len(vals))
			copy(buf2, vals)
			arr2 := append(buf2[:i+1], buf2[i+2:]...)
			if validateAscNoIssues(arr1) || validateAscNoIssues(arr2) {
				return true
			}
		}
	}

	return validateAscNoIssues(vals)
}

func validateAscNoIssues(vals []int) bool {
	for i := range vals {
		if i == len(vals)-1 {
			break
		}

		diff := vals[i+1] - vals[i]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func part1() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	result := 0
LineLoop:
	for scanner.Scan() {
		line := scanner.Text()

		strs := strings.Split(line, " ")
		if len(strs) < 2 {
			panic("incorect len")
		}
		vals := make([]int, len(strs))
		for i, s := range strs {
			val, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			vals[i] = val
		}

		if vals[1] == vals[0] {
			continue
		}

		isIncr, isDesc := false, false
		if vals[1]-vals[0] > 0 {
			isIncr = true
		} else {
			isDesc = true
		}

		for i := range vals {
			if i == len(vals)-1 {
				break
			}

			diff := vals[i+1] - vals[i]
			if isIncr {
				if diff < 1 || diff > 3 {
					continue LineLoop
				}
			}

			if isDesc {
				diff = vals[i] - vals[i+1]
				if diff < 1 || diff > 3 {
					continue LineLoop
				}
			}

		}

		result++
	}

	fmt.Println("Result: ", result)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
