package main

import (
	"bufio"
	"fmt"
	"os"
)

type C struct {
	X, Y int
}

var dirs = []C{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func main() {
	empties := make(map[C]bool)
	var e, s C
	for y, scanner := 0, bufio.NewScanner(os.Stdin); scanner.Scan(); y++ {
		for x, ch := range scanner.Text() {
			c := C{x, y}
			switch ch {
			case '.':
				empties[c] = true
			case 'S':
				s = c
				empties[c] = true
			case 'E':
				e = c
				empties[c] = true
			}
		}
	}

	thePath := []C{s}
	for thePath[len(thePath)-1] != e {
		for _, d := range dirs {
			n := C{thePath[len(thePath)-1].X + d.X, thePath[len(thePath)-1].Y + d.Y}
			if len(thePath) > 1 {
				if thePath[len(thePath)-2] == n {
					continue
				}
			}
			if empties[n] {
				thePath = append(thePath, n)
				break
			}
		}
	}

	distances := make(map[C]int)
	for i := 0; i < len(thePath); i++ {
		for j := len(thePath) - 1; j >= i+100+1; j-- {
			dist := abs(thePath[i].X-thePath[j].X) + abs(thePath[i].Y-thePath[j].Y)
			if dist >= 2 && dist <= 20 {
				saved := (j - i) - dist
				if saved >= 100 {
					distances[C{i, j}] = dist
				}
			}
		}
	}
	fmt.Println(len(distances))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
