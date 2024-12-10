package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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

	rulesCollected := false
	rulesBack := map[int][]int{} //ugh
	rulesFront := map[int][]int{}

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		// ugh
		if len(line) < 2 {
			rulesCollected = true
			continue
		}

		if !rulesCollected {
			strVals := strings.Split(line, "|")
			if len(strVals) != 2 {
				panic("incorrect len")
			}

			vals := []int{}
			for _, sv := range strVals {
				v, err := strconv.Atoi(sv)
				if err != nil {
					panic(fmt.Sprintf("incorrect val: %v\n", sv))
				}
				vals = append(vals, v)
			}

			first, second := vals[0], vals[1]
			rulesBack[first] = append(rulesBack[first], second)
			rulesFront[second] = append(rulesFront[second], first)
			continue
		}

		// rules collected
		strVals := strings.Split(line, ",")
		vals := []int{}
		for _, sv := range strVals {
			v, err := strconv.Atoi(sv)
			if err != nil {
				panic(fmt.Sprintf("incorrect val: %v\n", sv))
			}

			vals = append(vals, v)
		}

		// every value has to be good
		isGood := true
		for i, v := range vals {
			if prevVals, ok := rulesBack[v]; ok {
				if nextVals, ok := rulesFront[v]; ok {
					if !checkVals(vals, prevVals, nextVals, i) {
						isGood = false
						break
					}
				}
			}
		}

		if !isGood {
			sort.Slice(vals, func(i, j int) bool {
				first, second := vals[i], vals[j]
				if fPrevVals, ok := rulesBack[first]; ok {
					if slices.Contains(fPrevVals, second) {
						return true
					}
				}
				if sPrevVals, ok := rulesBack[second]; ok {
					if slices.Contains(sPrevVals, first) {
						return false
					}
				}

				if fNextVals, ok := rulesFront[first]; ok {
					if slices.Contains(fNextVals, second) {
						return false
					}
				}

				if sNextVals, ok := rulesFront[second]; ok {
					if slices.Contains(sNextVals, first) {
						return true
					}
				}
				return true
			})

			goodVar := true
			for i, v := range vals {
				if prevVals, ok := rulesBack[v]; ok {
					if nextVals, ok := rulesFront[v]; ok {
						if !checkVals(vals, prevVals, nextVals, i) {
							goodVar = false
							break
						}
					}
				}
			}
			if goodVar {
				result += vals[len(vals)/2]
			}

		}

	}

	fmt.Println("Result:", result)
}

func part1() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	rulesCollected := false
	rulesBack := map[int][]int{} //ugh
	rulesFront := map[int][]int{}

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		// ugh
		if len(line) < 2 {
			rulesCollected = true
			continue
		}

		if !rulesCollected {
			strVals := strings.Split(line, "|")
			if len(strVals) != 2 {
				panic("incorrect len")
			}

			vals := []int{}
			for _, sv := range strVals {
				v, err := strconv.Atoi(sv)
				if err != nil {
					panic(fmt.Sprintf("incorrect val: %v\n", sv))
				}
				vals = append(vals, v)
			}

			first, second := vals[0], vals[1]
			rulesBack[first] = append(rulesBack[first], second)
			rulesFront[second] = append(rulesFront[second], first)
			continue
		}

		// rules collected
		strVals := strings.Split(line, ",")
		vals := []int{}
		for _, sv := range strVals {
			v, err := strconv.Atoi(sv)
			if err != nil {
				panic(fmt.Sprintf("incorrect val: %v\n", sv))
			}

			vals = append(vals, v)
		}

		// every value has to be good
		isGood := true
		for i, v := range vals {
			if prevVals, ok := rulesBack[v]; ok {
				if nextVals, ok := rulesFront[v]; ok {
					if !checkVals(vals, prevVals, nextVals, i) {
						isGood = false
						break
					}
				}
			}
		}

		if isGood {
			result += vals[len(vals)/2]
		}
	}

	fmt.Println("Result:", result)
}

func checkVals(input, rulesBack, rulesFront []int, idx int) bool {
	// if there is a single value from rules behind input[idx] - false

	for _, v := range rulesBack {
		for i := 0; i < idx; i++ {
			if v == input[i] {
				return false
			}
		}
	}

	for _, v := range rulesFront {
		for i := idx + 1; i < len(input); i++ {
			if v == input[i] {
				return false
			}
		}
	}

	return true
}
