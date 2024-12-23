package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type C struct {
	A, B string
}

func main() {
	vs := make(map[string]bool)
	es := make(map[string]bool)
	candidates := make(map[C]bool)
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		bits := strings.Split(scanner.Text(), "-")
		slices.Sort(bits)
		es[strings.Join(bits, "-")] = true
		vs[bits[0]] = true
		vs[bits[1]] = true
		if bits[0][0] == 't' || bits[1][0] == 't' {
			candidates[C{bits[0], bits[1]}] = true
		}
	}

	finals := make(map[string]bool)
	for candidate := range candidates {
		// need to find a third
		for third := range vs {
			if third != candidate.A && third != candidate.B {
				oeA, oeB := []string{third, candidate.A}, []string{third, candidate.B}
				slices.Sort(oeA)
				slices.Sort(oeB)
				if es[strings.Join(oeA, "-")] && es[strings.Join(oeB, "-")] {

					fs := []string{third, candidate.A, candidate.B}
					slices.Sort(fs)
					finals[strings.Join(fs, "-")] = true
				}
			}
		}
	}
	fmt.Println(len(finals))
}
