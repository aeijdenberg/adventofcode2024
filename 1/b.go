package main

import (
	"bufio"
	"fmt"
	"os"
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
	c := make(map[int]int)
	for _, r := range right {
		c[r] += 1
	}
	s := 0
	for _, l := range left {
		s += l * c[l]
	}
	fmt.Println(s)
}
