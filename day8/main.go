package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// part1()
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

	// find all antennas
	// store them in a map to not bruteforce the whole grid
	antennas := map[byte][][2]int{}

	for i := range grid {
		for j := range grid[i] {
			cur := grid[i][j]
			if cur != '.' {
				antennas[cur] = append(antennas[cur], [2]int{i, j})
			}
		}
	}

	// get diff between every antenna of the same type
	// and try to place '#'
	// then count all the '#'

	locations := map[[2]int]bool{}
	for _, ant := range antennas {
		for i := range ant {
			for j := range ant {
				if i == j {
					continue
				}
				first := ant[i]
				second := ant[j]
				locations[first] = true
				locations[second] = true
				diff := [2]int{first[0] - second[0], first[1] - second[1]}

				newCoord := [2]int{first[0] + diff[0], first[1] + diff[1]}
				for isInBounds(newCoord, grid) {
					grid[newCoord[0]][newCoord[1]] = '#'
					locations[[2]int{newCoord[0], newCoord[1]}] = true
					newCoord[0] += diff[0]
					newCoord[1] += diff[1]
				}

			}
		}
	}

	fmt.Println("Result:", len(locations))
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

	// find all antennas
	// store them in a map to not bruteforce the whole grid
	antennas := map[byte][][2]int{}

	for i := range grid {
		for j := range grid[i] {
			cur := grid[i][j]
			if cur != '.' {
				antennas[cur] = append(antennas[cur], [2]int{i, j})
			}
		}
	}

	// get diff between every antenna of the same type
	// and try to place '#'
	// then count all the '#'

	locations := map[[2]int]bool{}
	for _, ant := range antennas {
		for i := range ant {
			for j := range ant {
				if i == j {
					continue
				}
				first := ant[i]
				second := ant[j]
				diff := [2]int{first[0] - second[0], first[1] - second[1]}

				newCoord := [2]int{first[0] + diff[0], first[1] + diff[1]}
				if isInBounds(newCoord, grid) {
					// grid[newCoord[0]][newCoord[1]] = '#'
					locations[[2]int{newCoord[0], newCoord[1]}] = true
				}

			}
		}
	}

	fmt.Println("Result:", len(locations))
}

func isInBounds(coord [2]int, grid [][]byte) bool {
	i := coord[0]
	j := coord[1]
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
