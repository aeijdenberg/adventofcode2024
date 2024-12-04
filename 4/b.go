package main

import (
	"bufio"
	"fmt"
	"os"
)

func isCh(lines [][]byte, ch byte, x, y int) bool {
	if y < 0 || y >= len(lines) {
		return false
	}
	line := lines[y]
	if x < 0 || x >= len(line) {
		return false
	}
	return lines[y][x] == ch
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines [][]byte
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	search := []byte("MAS")

	spots := make(map[struct{ x, y int }]int)
	for y, line := range lines {
		for x := range line {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 || dy == 0 {
						continue
					}
					matches := true
					for i, ch := range search {
						if !isCh(lines, ch, x+(i*dx), y+(i*dy)) {
							matches = false
						}
					}
					if matches {
						spots[struct{ x, y int }{x + dx, y + dy}]++
					}
				}
			}
		}
	}
	s := 0
	for _, count := range spots {
		if count == 2 {
			s++
		}
	}

	fmt.Println(s)
}
