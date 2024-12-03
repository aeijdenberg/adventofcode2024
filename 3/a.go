package main

import (
	"fmt"
	"io"
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

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	s := 0
	for _, results := range regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(string(b), -1) {
		s += MustInt(results[1]) * MustInt(results[2])
	}

	fmt.Println(s)
}
