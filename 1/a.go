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
	var left, right []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bits := strings.Split(scanner.Text(), "   ")
		left = append(left, MustInt(bits[0]))
		right = append(right, MustInt(bits[1]))
	}
	slices.Sort(left)
	slices.Sort(right)
	s := 0
	for i, l := range left {
		d := right[i] - l
		if d < 0 {
			d *= -1
		}
		s += d
	}
	fmt.Println(s)
}
