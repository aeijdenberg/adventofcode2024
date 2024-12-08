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

func abs(d int) int {
	if d >= 0 {
		return d
	}
	return -d
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
			for z := -w; z <= w; z++ { // use w as arbitrary TOO many indicator, e.g. we don't care if out of bounds
				antinodes[C{c1.X - z*dx, c1.Y - z*dy}] = true
				antinodes[C{c2.X + z*dx, c2.Y + z*dy}] = true
			}
		})
	}

	s := 0
	for c := range antinodes {
		if c.X >= 0 && c.X < w && c.Y >= 0 && c.Y < h {
			s++
		}
	}

	fmt.Println(s, len(antinodes))
}
