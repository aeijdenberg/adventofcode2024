package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
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

type CW struct {
	C int
}

func (c *CW) Write(b []byte) (int, error) {
	c.C += len(b)
	return len(b), nil
}

var cw = &CW{}
var gw = gzip.NewWriter(cw)

func compressSize(b []byte) int {
	cw.C = 0
	gw.Reset(cw)
	_, err := gw.Write(b)
	if err != nil {
		panic(err)
	}
	err = gw.Close()
	if err != nil {
		panic(err)
	}
	return cw.C
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
	buf := &bytes.Buffer{}
	best := 0x3fffffff
	for i := 0; ; i++ {
		counts := make(map[C]int)

		for _, r := range rs {
			counts[C{mod(r.P.X+r.V.X*i, w), mod(r.P.Y+r.V.Y*i, h)}] += 1
		}

		buf.Reset()
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				c := C{x, y}
				if counts[c] == 0 {
					fmt.Fprintf(buf, ".")
				} else {
					fmt.Fprintf(buf, "X")
				}
			}
			fmt.Fprintf(buf, "\n")
		}

		score := compressSize(buf.Bytes())
		if score < best {
			fmt.Printf("%d (score: %d):\n%s\n", i, score, buf.Bytes())
			best = score
		}
	}

}
