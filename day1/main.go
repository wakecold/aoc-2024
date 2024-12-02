package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//part1()
	part2()
}

func part2() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	uniqL := map[int]bool{}
	countR := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.Split(line, " ")

		first, second := vals[0], vals[len(vals)-1]
		l, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(second)
		if err != nil {
			panic(err)
		}

		uniqL[l] = true
		countR[r]++
	}

	var result int64

	for val := range uniqL {
		if count, ok := countR[val]; ok {
			result += int64(val * count)
		}
	}

	fmt.Println("Result: ", result)
}

func part1() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	left, right := []int{}, []int{}

	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.Split(line, " ")

		first, second := vals[0], vals[len(vals)-1]
		l, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(second)
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	result := 0
	for i := range left {
		result += abs(right[i] - left[i])
	}

	fmt.Println("result:", result)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
