package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Line struct {
	Start int
	Rise  int
}

func (l Line) String() string {
	return fmt.Sprintf("%dx + %d", l.Rise, l.Start)
}

func (l *Line) Vals() chan int {
	rv := make(chan int)
	go func() {
		defer close(rv)
		if l == nil {
			return
		}
		n := l.Start
		for {
			rv <- n
			if l.Rise == 0 {
				return
			}
			n += l.Rise
		}
	}()
	return rv
}

func findMatches(c1, c2 chan int, attempts int) chan int {
	rv := make(chan int)
	go func() {
		defer close(rv)

		v1, ok := <-c1
		if !ok {
			return
		}

		v2, ok := <-c2
		if !ok {
			return
		}

		for {
			if attempts == 0 {
				return
			}
			attempts--
			for v1 < v2 {
				v1, ok = <-c1
				if !ok {
					return
				}
			}
			for v2 < v1 {
				v2, ok = <-c2
				if !ok {
					return
				}
			}
			if v1 == v2 {
				rv <- v1
				v1, ok = <-c1
				if !ok {
					return
				}
			}
		}
	}()
	return rv
}

func generateCleanFs(ax, bx, px, attempts int) chan int {
	rv := make(chan int)
	go func() {
		defer close(rv)
		for f := 0; f*ax < px && attempts > 0; f++ {
			if (px-f*ax)%bx == 0 {
				rv <- f
			}
			attempts--
		}
	}()
	return rv
}

func makeLine(vals chan int) *Line {
	var rv Line
	var ok bool
	rv.Start, ok = <-vals
	if !ok {
		return nil
	}
	v, ok := <-vals
	if !ok {
		return &rv
	}
	rv.Rise = v - rv.Start
	return &rv
}

func solve(ax, ay, bx, by, px, py int) int {
	// I have no idea why this works
	fCands := makeLine(findMatches(makeLine(generateCleanFs(ax, bx, px, 200)).Vals(), makeLine(generateCleanFs(ay, by, py, 200)).Vals(), 200))

	var lastD int
	first := true
	vals := fCands.Vals()
	for {
		f, ok := <-vals
		if !ok {
			return 0
		}

		g1, g2 := (px-f*ax)/bx, (py-f*ay)/by
		if g1 == g2 {
			return f*3 + g1
		}

		delta := g1 - g2
		if first {
			first = false
		} else {
			if delta%(lastD-delta) != 0 {
				return 0
			}

			f = f + fCands.Rise*(delta/(lastD-delta))

			g1, g2 := (px-f*ax)/bx, (py-f*ay)/by
			if g1 == g2 {
				return f*3 + g1
			} else {
				panic("bad")
			}
		}

		lastD = delta
	}
}

func MustInt(s string) int {
	rv, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return rv
}

func main() {
	r := regexp.MustCompile(`X[=+](\d+), Y[=+](\d+)`)
	var ax, ay, bx, by, px, py, s int
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "Button A:"):
			bits := r.FindStringSubmatch(line)
			ax, ay = MustInt(bits[1]), MustInt(bits[2])

		case strings.HasPrefix(line, "Button B:"):
			bits := r.FindStringSubmatch(line)
			bx, by = MustInt(bits[1]), MustInt(bits[2])

		case strings.HasPrefix(line, "Prize:"):
			bits := r.FindStringSubmatch(line)
			px, py = MustInt(bits[1]), MustInt(bits[2])

			px += 10000000000000
			py += 10000000000000

			s += solve(ax, ay, bx, by, px, py)
		}

	}
	fmt.Println(s)
}
