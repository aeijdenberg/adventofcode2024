package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type C struct {
	X, Y int
}

type Keypad map[SE][]string

type SE struct {
	S, E byte
}

func NewKeypad(m map[byte]C) Keypad {
	im := make(map[C]byte)
	for k, v := range m {
		im[v] = k
	}
	rv := make(Keypad)
	for s := range m {
		for e := range m {
			rv[SE{s, e}] = enumeratePaths(s, e, m, im)
		}
	}
	return rv
}

func tryThese(prefix string, sc C, e byte, m map[byte]C, im map[C]byte) []string {
	s, ok := im[sc]
	if !ok {
		return nil
	}
	rv := enumeratePaths(s, e, m, im)
	for i, p := range rv {
		rv[i] = prefix + p
	}
	return rv
}

func enumeratePaths(s, e byte, m map[byte]C, im map[C]byte) []string {
	if s == e {
		return []string{"A"}
	}
	sc, ec := m[s], m[e]
	if sc.X == ec.X {
		if sc.Y < ec.Y {
			return tryThese("v", C{sc.X, sc.Y + 1}, e, m, im)
		}
		if sc.Y > ec.Y {
			return tryThese("^", C{sc.X, sc.Y - 1}, e, m, im)
		}
		panic("oh noe")
	}
	if sc.Y == ec.Y {
		if sc.X < ec.X {
			return tryThese(">", C{sc.X + 1, sc.Y}, e, m, im)
		}
		if sc.X > ec.X {
			return tryThese("<", C{sc.X - 1, sc.Y}, e, m, im)
		}
		panic("oh noe")
	}
	// both are different
	if sc.X < ec.X && sc.Y < ec.Y {
		return append(tryThese(">", C{sc.X + 1, sc.Y}, e, m, im), tryThese("v", C{sc.X, sc.Y + 1}, e, m, im)...)
	}
	if sc.X > ec.X && sc.Y < ec.Y {
		return append(tryThese("<", C{sc.X - 1, sc.Y}, e, m, im), tryThese("v", C{sc.X, sc.Y + 1}, e, m, im)...)
	}
	if sc.X < ec.X && sc.Y > ec.Y {
		return append(tryThese(">", C{sc.X + 1, sc.Y}, e, m, im), tryThese("^", C{sc.X, sc.Y - 1}, e, m, im)...)
	}
	if sc.X > ec.X && sc.Y > ec.Y {
		return append(tryThese("<", C{sc.X - 1, sc.Y}, e, m, im), tryThese("^", C{sc.X, sc.Y - 1}, e, m, im)...)
	}
	panic("oh noes")
}

var (
	Numeric = NewKeypad(map[byte]C{
		'7': {0, 0},
		'8': {1, 0},
		'9': {2, 0},
		'4': {0, 1},
		'5': {1, 1},
		'6': {2, 1},
		'1': {0, 2},
		'2': {1, 2},
		'3': {2, 2},
		'0': {1, 3},
		'A': {2, 3},
	})
	Directional = NewKeypad(map[byte]C{
		'^': {1, 0},
		'A': {2, 0},
		'<': {0, 1},
		'v': {1, 1},
		'>': {2, 1},
	})
)

func (k Keypad) CountButtons(code string, scorePath func(p string) int) int {
	last := byte('A')
	rv := 0
	for _, ch := range []byte(code) {
		paths := k[SE{last, ch}]
		if len(paths) == 0 {
			panic("wrong")
		}
		best := math.MaxInt
		for _, p := range paths {
			best = min(best, scorePath(p))
		}
		rv += best
		last = ch
	}
	return rv
}

func solve(code string) int {
	return Numeric.CountButtons(code, func(p string) int {
		return Directional.CountButtons(p, func(p string) int {
			return Directional.CountButtons(p, func(p string) int { return len(p) })
		})
	})
}

func MustInt(s string) int {
	rv, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return rv
}

func main() {
	s := 0
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()
		s += MustInt(regexp.MustCompile(`[1-9][0-9]*`).FindString(line)) * solve(line)
	}
	fmt.Println(s)
}
