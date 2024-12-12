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
		lefts, rights, ups, downs := make(map[C]bool), make(map[C]bool), make(map[C]bool), make(map[C]bool)
		for c := range r {
			if !r[C{c.X, c.Y - 1}] {
				ups[C{c.X, c.Y - 1}] = true
			}
			if !r[C{c.X, c.Y + 1}] {
				downs[C{c.X, c.Y + 1}] = true
			}
			if !r[C{c.X - 1, c.Y}] {
				lefts[C{c.X - 1, c.Y}] = true
			}
			if !r[C{c.X + 1, c.Y}] {
				rights[C{c.X + 1, c.Y}] = true
			}
		}

		sides := 0

		done := make(map[C]bool)
		for c := range ups {
			if !done[c] {
				sides++
				done[c] = true
				for i := 1; ups[C{c.X - i, c.Y}]; i++ {
					done[C{c.X - i, c.Y}] = true
				}
				for i := 1; ups[C{c.X + i, c.Y}]; i++ {
					done[C{c.X + i, c.Y}] = true
				}
			}
		}

		done = make(map[C]bool)
		for c := range downs {
			if !done[c] {
				sides++
				done[c] = true
				for i := 1; downs[C{c.X - i, c.Y}]; i++ {
					done[C{c.X - i, c.Y}] = true
				}
				for i := 1; downs[C{c.X + i, c.Y}]; i++ {
					done[C{c.X + i, c.Y}] = true
				}
			}
		}

		done = make(map[C]bool)
		for c := range lefts {
			if !done[c] {
				sides++
				done[c] = true
				for i := 1; lefts[C{c.X, c.Y - i}]; i++ {
					done[C{c.X, c.Y - i}] = true
				}
				for i := 1; lefts[C{c.X, c.Y + i}]; i++ {
					done[C{c.X, c.Y + i}] = true
				}
			}
		}

		done = make(map[C]bool)
		for c := range rights {
			if !done[c] {
				sides++
				done[c] = true
				for i := 1; rights[C{c.X, c.Y - i}]; i++ {
					done[C{c.X, c.Y - i}] = true
				}
				for i := 1; rights[C{c.X, c.Y + i}]; i++ {
					done[C{c.X, c.Y + i}] = true
				}
			}
		}

		s += sides * len(r)
	}
	fmt.Println(s)
}
