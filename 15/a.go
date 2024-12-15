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
					m[C{x, y}] = byte(ch)
					if ch == '@' {
						r = C{x, y}
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
		} else if m[next] == 'O' {
			last := next
			for m[last] == 'O' {
				last = C{last.X + d.X, last.Y + d.Y}
			}
			if m[last] == '.' {
				m[last] = 'O'
				m[next] = '@'
				m[r] = '.'
				r = next
			}
		}
	}

	s := 0
	for c, t := range m {
		if t == 'O' {
			s += 100*c.Y + c.X
		}
	}

	fmt.Println(s)
}
