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

	// get the route, try each location for a new wall
	// if visited > 4 then its a loop ??
	startI, startJ := 0, 0
	route := [][]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "^" {
				startI, startJ = i, j
				findRoute(i, j, grid, []int{-1, 0}, &route)
			}
		}
	}

	// remove starting pos?
	route = route[1:]

	newWalls := map[[2]int]bool{}
	for _, r := range route {
		grid[r[0]][r[1]] = "#"
		if isLoop(startI, startJ, grid, []int{-1, 0}, map[[2]int]int{}) {
			newWalls[[2]int{r[0], r[1]}] = true
		}

		grid[r[0]][r[1]] = "."
	}

	fmt.Println("Result:", len(newWalls))
}

func isLoop(i, j int, grid [][]string, dir []int, freq map[[2]int]int) bool {
	if c, ok := freq[[2]int{i, j}]; ok {
		if c > 4 {
			return true
		}
	}

	for i >= 0 &&
		j >= 0 &&
		i < len(grid) &&
		j < len(grid[0]) &&
		grid[i][j] != "#" {
		freq[[2]int{i, j}]++
		i += dir[0]
		j += dir[1]
	}

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return false
	}

	// step back
	i -= dir[0]
	j -= dir[1]

	nextDir := []int{dir[1], -dir[0]}

	return isLoop(i, j, grid, nextDir, freq)
}

func findRoute(i, j int, grid [][]string, dir []int, route *[][]int) {
	for i >= 0 &&
		j >= 0 &&
		i < len(grid) &&
		j < len(grid[0]) &&
		grid[i][j] != "#" {
		(*route) = append(*route, []int{i, j})
		i += dir[0]
		j += dir[1]
	}

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	(*route) = (*route)[:len(*route)-1]
	// step back
	i -= dir[0]
	j -= dir[1]

	nextDir := []int{dir[1], -dir[0]}

	findRoute(i, j, grid, nextDir, route)
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

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "^" {
				tr(i, j, grid, []int{-1, 0})
			}
		}
	}

	result := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "X" {
				result++
			}
		}
	}

	fmt.Println("Result:", result)
}

func tr(i, j int, grid [][]string, dir []int) {
	for i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0]) && grid[i][j] != "#" {
		grid[i][j] = "X"
		i += dir[0]
		j += dir[1]
	}

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	// step back
	i -= dir[0]
	j -= dir[1]

	nextDir := []int{dir[1], -dir[0]}

	tr(i, j, grid, nextDir)
}
