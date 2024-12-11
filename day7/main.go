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

		strVals := strings.Split(line, ":")
		if len(strVals) != 2 {
			panic("incorrect string")
		}

		lineResult, err := strconv.Atoi(strVals[0])
		if err != nil {
			panic("incorrect value")
		}

		ops := strings.Split(strVals[1], " ")
		vals := []int{}
		for _, op := range ops {
			val, err := strconv.Atoi(op)
			if err != nil {
				// skip
				continue
			}
			vals = append(vals, val)
		}

		if len(vals) == 0 {
			panic("empty vals")
		}

		if dpWithConcat(1, vals, int64(vals[0]), lineResult) {
			result += lineResult
		}
	}

	fmt.Println("Result:", result)
}

func part1() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()

		strVals := strings.Split(line, ":")
		if len(strVals) != 2 {
			panic("incorrect string")
		}

		lineResult, err := strconv.Atoi(strVals[0])
		if err != nil {
			panic("incorrect value")
		}

		ops := strings.Split(strVals[1], " ")
		vals := []int{}
		for _, op := range ops {
			val, err := strconv.Atoi(op)
			if err != nil {
				// skip
				continue
			}
			vals = append(vals, val)
		}

		if len(vals) == 0 {
			panic("empty vals")
		}

		if dp(1, vals, vals[0], lineResult) {
			result += lineResult
		}
	}

	fmt.Println("Result:", result)
}

func dp(i int, vals []int, cur int, result int) bool {
	if i >= len(vals) {
		return cur == result
	}

	return dp(i+1, vals, cur+vals[i], result) || dp(i+1, vals, cur*vals[i], result)
}

func dpWithConcat(i int, vals []int, cur int64, result int) bool {
	if cur > int64(result) {
		return false
	}
	if i >= len(vals) {
		return cur == int64(result)
	}

	curStr := strconv.FormatInt(cur, 10)
	idxValStr := strconv.FormatInt(int64(vals[i]), 10)
	concatStr := curStr + idxValStr
	concatVal, err := strconv.ParseInt(concatStr, 10, 64)
	if err != nil {
		fmt.Println(concatStr)
		panic("bad concat value")
	}

	sumVal := cur + int64(vals[i])
	multVal := cur * int64(vals[i])

	return dpWithConcat(i+1, vals, sumVal, result) ||
		dpWithConcat(i+1, vals, multVal, result) ||
		dpWithConcat(i+1, vals, concatVal, result)

}
