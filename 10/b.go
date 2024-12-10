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
	{-1, 0},
	{0, 1},
	{0, -1},
}

func main() {
	m := make(map[C]int)
	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		for x, ch := range scanner.Bytes() {
			m[C{x, y}] = int(ch - '0')
		}
	}

	cache := make(map[C]int)
	s := 0
	for i := 9; i >= 0; i-- {
		for k, v := range m {
			if v == i {
				if i == 9 {
					cache[k] = 1
				} else {
					// find all adjacent, and add their things to us
					for _, d := range dirs {
						n := C{k.X + d.X, k.Y + d.Y}
						if m[n] == i+1 {
							cache[k] += cache[n]
						}
					}

					if i == 0 {
						s += cache[k]
					}
				}
			}
		}
	}

	fmt.Println(s)
}
