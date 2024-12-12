package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = []C{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type C struct {
	X, Y int
}

func addToRegion(m map[C]byte, r, d map[C]bool, c C, v byte) {
	if d[c] {
		return
	}
	r[c] = true
	d[c] = true
	for _, dd := range dirs {
		c2 := C{c.X + dd.X, c.Y + dd.Y}
		if m[c2] == v {
			addToRegion(m, r, d, c2, v)
		}
	}
}

func main() {
	m := make(map[C]byte)
	for s, y := bufio.NewScanner(os.Stdin), 0; s.Scan(); y++ {
		for x, ch := range []byte(s.Text()) {
			m[C{x, y}] = ch
		}
	}
	var regions []map[C]bool
	done := make(map[C]bool)
	for c, v := range m {
		if !done[c] {
			r := make(map[C]bool)
			addToRegion(m, r, done, c, v)
			regions = append(regions, r)
		}
	}

	s := 0
	for _, r := range regions {
		p := 0
		for c := range r {
			for _, dd := range dirs {
				c2 := C{c.X + dd.X, c.Y + dd.Y}
				if !r[c2] {
					p++
				}
			}
		}
		s += p * len(r)
	}
	fmt.Println(s)
}
