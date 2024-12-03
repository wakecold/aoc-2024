package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	result := 0
	line := ""
	for scanner.Scan() {
		line += scanner.Text()
	}

	i := 0
	mulEnabled := true
LineLoop:
	for i < len(line) {
		if i < len(line) &&
			line[i] == 'd' &&
			i+3 < len(line) &&
			line[i+1] == 'o' &&
			line[i+2] == '(' &&
			line[i+3] == ')' {
			mulEnabled = true
			i = i + 4
		}
		if i < len(line) &&
			line[i] == 'd' &&
			i+6 < len(line) &&
			line[i+1] == 'o' &&
			line[i+2] == 'n' &&
			line[i+3] == '\'' &&
			line[i+4] == 't' &&
			line[i+5] == '(' &&
			line[i+6] == ')' {
			mulEnabled = false
			i = i + 7
		}
		if i < len(line) && line[i] == 'm' && i+2 < len(line) && line[i+1] == 'u' && line[i+2] == 'l' {
			i = i + 3

			if i < len(line) && line[i] == '(' {
				s := ""
				next := i + 1
				for line[next] != ')' && next < len(line) {
					// special case
					if line[next] == 'm' {
						i = next
						continue LineLoop
					}
					if line[next] == 'd' {
						i = next
						continue LineLoop
					}
					s += string(line[next])
					next++
				}

				strs := strings.Split(s, ",")
				if len(strs) != 2 {
					i = next
					continue
				}
				left, err := strconv.Atoi(strs[0])
				if err != nil {
					i = next
					continue
				}

				right, err := strconv.Atoi(strs[1])
				if err != nil {
					i = next
					continue
				}

				if mulEnabled {
					result += left * right
				}
			}
		}
		i++
	}

	fmt.Println("Result:", result)
}
