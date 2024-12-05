package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func MustInt(s string) int {
	rv, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return rv
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	state := 0
	lower := make(map[string]bool)
	rv := 0
	for scanner.Scan() {
		s := scanner.Text()
		switch state {
		case 0:
			if len(s) == 0 {
				state = 1
				continue
			}
			lower[s] = true
		case 1:
			bits := strings.Split(s, ",")
			isGood := true
			for a := 0; isGood && a < len(bits)-1; a++ {
				for b := a + 1; isGood && b < len(bits); b++ {
					if lower[bits[b]+"|"+bits[a]] {
						isGood = false
					}
				}
			}
			if !isGood {
				slices.SortStableFunc(bits, func(a, b string) int {
					if lower[a+"|"+b] {
						return -1
					}
					return 0
				})
				rv += MustInt(bits[(len(bits)-1)/2])
			}
		}
	}
	fmt.Println(rv)
}
