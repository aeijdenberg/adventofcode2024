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
	w, h := 101, 103
	for i := 0; ; i++ {
		counts := make(map[C]int)

		for _, r := range rs {
			counts[C{mod(r.P.X+r.V.X*i, w), mod(r.P.Y+r.V.Y*i, h)}] += 1
		}

		if len(counts) != len(rs) {
			continue
		}

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				c := C{x, y}
				if counts[c] == 0 {
					fmt.Printf(".")
				} else {
					fmt.Printf("%d", counts[c])
				}
			}
			fmt.Printf("\n")
		}
		fmt.Println(i)
		break
	}

}
