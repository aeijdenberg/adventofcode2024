package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func MustInt(s string) int {
	rv, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return rv
}

type C struct {
	X, Y int
}

type R struct {
	P, V C
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	r := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	var rs []*R
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		bits := r.FindStringSubmatch(scanner.Text())
		rs = append(rs, &R{
			P: C{X: MustInt(bits[1]), Y: MustInt(bits[2])},
			V: C{X: MustInt(bits[3]), Y: MustInt(bits[4])},
		})
	}
	w, h, s := 101, 103, 100
	qs := make(map[string]int)
	for _, r := range rs {
		r.P.X = mod(r.P.X+r.V.X*s, w)
		r.P.Y = mod(r.P.Y+r.V.Y*s, h)

		xq := ""
		if r.P.X < ((w - 1) / 2) {
			xq = "L"
		} else if r.P.X > ((w - 1) / 2) {
			xq = "R"
		}
		yq := ""
		if r.P.Y < ((h - 1) / 2) {
			yq = "T"
		} else if r.P.Y > ((h - 1) / 2) {
			yq = "B"
		}

		if xq != "" && yq != "" {
			q := xq + yq
			qs[q] = qs[q] + 1
		}

	}

	rv := 1
	for _, c := range qs {
		rv *= c
	}

	fmt.Println(rv)
}
