package main

import (
	"bufio"
	"fmt"
	"os"
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

	grid := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.Split(line, "")

		grid = append(grid, vals)
	}

	// Find all A and go all directions to check if its a cross

	result := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "A" {
				result += checkCross(i, j, grid)
			}
		}
	}

	fmt.Println("Result:", result)
}

func checkCross(i, j int, grid [][]string) int {
	result := 0

	if isInBound(i-1, j-1, grid) &&
		isInBound(i+1, j+1, grid) &&
		((grid[i-1][j-1] == "M" &&
			grid[i+1][j+1] == "S") ||
			(grid[i-1][j-1] == "S") &&
				(grid[i+1][j+1] == "M")) &&
		(isInBound(i+1, j-1, grid) &&
			isInBound(i-1, j+1, grid) &&
			((grid[i+1][j-1] == "M" &&
				grid[i-1][j+1] == "S") ||
				(grid[i+1][j-1] == "S" &&
					grid[i-1][j+1] == "M"))) {
		result++
	}

	return result
}

func isInBound(i, j int, grid [][]string) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func part1() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	grid := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.Split(line, "")

		grid = append(grid, vals)
	}

	result := 0
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "X" {
				for _, dir := range dirs {
					res := dfs(i+dir[0], j+dir[1], grid, "MAS", dir)
					if res {
						result++
					}
				}
			}
		}
	}

	fmt.Println("Result:", result)

}

func dfs(i, j int, grid [][]string, s string, dir []int) bool {
	if len(s) == 0 {
		return true
	}
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return false
	}

	if grid[i][j] != string(s[0]) {
		return false
	}

	return dfs(i+dir[0], j+dir[1], grid, s[1:], dir)
}
