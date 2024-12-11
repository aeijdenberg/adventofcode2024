package main

import (
	"fmt"
	"io"
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

type C struct {
	V, B int
}

func NumDigits(v int) int {
	s := strconv.Itoa(v)
	return len(s)
}

func solve(cache map[C]int, v, b int) int {
	if b == 0 {
		return 1
	}

	rv, ok := cache[C{v, b}]
	if ok {
		return rv
	}

	if v == 0 {
		rv = solve(cache, 1, b-1)
	} else {
		s := strconv.Itoa(v)
		if len(s)%2 == 0 {
			rv = solve(cache, MustInt(s[:len(s)/2]), b-1) + solve(cache, MustInt(s[len(s)/2:]), b-1)
		} else {
			rv = solve(cache, v*2024, b-1)
		}
	}

	cache[C{v, b}] = rv
	return rv
}

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	cache := make(map[C]int)
	s := 0
	for _, bit := range strings.Split(strings.TrimSpace(string(b)), " ") {
		s += solve(cache, MustInt(bit), 25)
	}
	fmt.Println(s)
}
