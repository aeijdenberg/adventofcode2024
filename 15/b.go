package main

import (
	"bufio"
	"fmt"
	"os"
)

type C struct {
	X, Y int
}

var dirs = map[byte]C{
	'<': {-1, 0},
	'>': {1, 0},
	'v': {0, 1},
	'^': {0, -1},
}

func resolveBox(c C, ch byte) C {
	if ch == '[' {
		return c
	} else {
		return C{c.X - 1, c.Y}
	}
}

func getMoveableBoxes(m map[C]byte, box C, dy int) []C {
	lC, rC := C{box.X, box.Y + dy}, C{box.X + 1, box.Y + dy}
	l, r := m[lC], m[rC]
	if l == '.' && r == '.' {
		return []C{box}
	}
	if l == '#' || r == '#' {
		return nil
	}
	// else both could be box fragments
	if l == '.' { // just the right is a box
		rv := getMoveableBoxes(m, resolveBox(rC, r), dy)
		if len(rv) == 0 {
			return nil
		}
		return append(rv, box)
	}

	if r == '.' { // just the left is a box
		rv := getMoveableBoxes(m, resolveBox(lC, l), dy)
		if len(rv) == 0 {
			return nil
		}
		return append(rv, box)
	}

	// both is boxes
	rvL := getMoveableBoxes(m, resolveBox(lC, l), dy)
	rvR := getMoveableBoxes(m, resolveBox(rC, r), dy)
	if len(rvL) == 0 || len(rvR) == 0 {
		return nil
	}
	return append(append(rvL, rvR...), box)
}

func main() {
	m := make(map[C]byte)
	var moves []byte
	var r C

	state := 0
	for y, scanner := 0, bufio.NewScanner(os.Stdin); scanner.Scan(); y++ {
		s := scanner.Text()
		switch state {
		case 0:
			if s == "" {
				state = 1
			} else {
				for x, ch := range s {
					switch ch {
					case '@':
						r = C{x * 2, y}
						m[C{x * 2, y}] = byte(ch)
						m[C{x*2 + 1, y}] = '.'
					case 'O':
						m[C{x * 2, y}] = '['
						m[C{x*2 + 1, y}] = ']'
					default:
						m[C{x * 2, y}] = byte(ch)
						m[C{x*2 + 1, y}] = byte(ch)
					}
				}
			}
		case 1:
			moves = append(moves, []byte(s)...)
		}
	}

	for _, mv := range moves {
		d := dirs[mv]
		next := C{r.X + d.X, r.Y + d.Y}
		if m[next] == '.' {
			m[next] = '@'
			m[r] = '.'
			r = next
		} else if m[next] == '[' || m[next] == ']' {
			if d.X == 0 {
				// vertical
				boxesToMove := getMoveableBoxes(m, resolveBox(next, m[next]), d.Y)
				if len(boxesToMove) != 0 {
					// first zero them all then
					// then put them all next spot
					for _, cb := range boxesToMove {
						m[cb] = '.'
						m[C{cb.X + 1, cb.Y}] = '.'
					}
					for _, cb := range boxesToMove {
						m[C{cb.X, cb.Y + d.Y}] = '['
						m[C{cb.X + 1, cb.Y + d.Y}] = ']'
					}
					// the move robot
					m[next] = '@'
					m[r] = '.'
					r = next
				}
			} else {
				// horizontal
				ogBox := m[next]

				last := next
				for m[last] == ogBox {
					last = C{last.X + 2*d.X, r.Y}
				}

				if m[C{last.X, r.Y}] == '.' {
					for x := last.X; x != r.X; x -= d.X {
						m[C{x, r.Y}] = m[C{x - d.X, r.Y}]
					}
					m[r] = '.'
					r = next
				}
			}
		}
	}

	for y := 0; y < 50; y++ {
		for x := 0; x < 100; x++ {
			fmt.Printf("%c", m[C{x, y}])
		}
		fmt.Printf("\n")
	}

	s := 0
	for c, t := range m {
		if t == '[' {
			s += 100*c.Y + c.X
		}
	}

	fmt.Println(s)
}
