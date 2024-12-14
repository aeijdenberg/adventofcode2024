package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func solve(ax, ay, bx, by, px, py int) int {
	f := (px*by - py*bx) / (ax*by - ay*bx)
	g := (px - f*ax) / bx
	if ax*f+bx*g != px || ay*f+by*g != py {
		return 0
	}
	return 3*f + g
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
	fmt.Println(s) // 89013607072065
}
