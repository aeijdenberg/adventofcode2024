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

	args := regexp.MustCompile(`([\d]+),([\d]+)`)

	s := 0
	enabled := true
	for _, results := range regexp.MustCompile(`(mul|do|don't)\((\d+,\d+)?\)`).FindAllStringSubmatch(string(b), -1) {
		fmt.Println(results)
		switch results[1] {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if !enabled {
				continue
			}
			for _, sr := range args.FindAllStringSubmatch(results[2], -1) {
				s += MustInt(sr[1]) * MustInt(sr[2])
			}
		}
	}

	fmt.Println(s)
}
