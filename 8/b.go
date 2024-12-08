package main

import (
	"bufio"
	"fmt"
	"os"
)

func ForEach(a []C, f func(i, j C)) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			f(a[i], a[j])
		}
	}
}

type C struct {
	X, Y int
}

func addIfInBounds(antinodes map[C]bool, w, h int, a C) bool {
	if a.X < 0 || a.X >= w || a.Y < 0 || a.Y >= h {
		return false
	}
	antinodes[a] = true
	return true
}

func or(a, b bool) bool {
	return a || b
}

func main() {
	var w, h int
	freqs := make(map[byte][]C)
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); h++ {
		w = 0
		for _, ch := range scanner.Bytes() {
			if ch != '.' {
				freqs[ch] = append(freqs[ch], C{w, h})
			}
			w++
		}
	}

	antinodes := make(map[C]bool)
	for _, cs := range freqs {
		ForEach(cs, func(c1, c2 C) {
			dx := c2.X - c1.X
			dy := c2.Y - c1.Y
			for z := 0; or(addIfInBounds(antinodes, w, h, C{c1.X - z*dx, c1.Y - z*dy}), addIfInBounds(antinodes, w, h, C{c2.X + z*dx, c2.Y + z*dy})); z++ {
			}
		})
	}

	fmt.Println(len(antinodes))
}
