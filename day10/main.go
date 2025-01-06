package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part2()
}

func part2() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	result := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				result += dfsPart2(i, j, 0, grid)
			}
		}
	}

	fmt.Println("Result:", result)
}

func dfsPart2(i, j int, cur int, grid [][]byte) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if cur == 9 {
		if grid[i][j] == '9' {
			return 1
		}
		return 0
	}
	if int(grid[i][j])-int('0') != cur {
		return 0
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	result := 0
	for _, dir := range dirs {
		result += dfsPart2(i+dir[0], j+dir[1], cur+1, grid)
	}
	return result
}

func part1() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	result := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				cur := map[[2]int]bool{}
				dfsPart1(i, j, 0, cur, grid)
				result += len(cur)
			}
		}
	}
	fmt.Println("Result:", result)
}

func dfsPart1(i, j int, cur int, res map[[2]int]bool, grid [][]byte) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if cur == 9 {
		if grid[i][j] == '9' {
			res[[2]int{i, j}] = true
			return
		}
		return
	}
	if int(grid[i][j])-int('0') != cur {
		return
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		dfsPart1(i+dir[0], j+dir[1], cur+1, res, grid)
	}
}
