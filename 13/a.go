package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solve(ax, ay, bx, by, px, py int) int {
	best := 1000
	for f := 0; f <= 100; f++ {
		for g := 0; g <= 100; g++ {
			potScore := f*3 + g
			if potScore < best {
				if (ax*f)+(bx*g) == px && (ay*f)+(by*g) == py {
					best = potScore
				}
			}
		}
	}
	if best == 1000 {
		return 0
	}
	return best
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

			s += solve(ax, ay, bx, by, px, py)
		}

	}
	fmt.Println(s)
}
