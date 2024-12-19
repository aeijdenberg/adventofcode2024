package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cache = make(map[string]int)

func possible(words map[string]bool, target string) int {
	if len(target) == 0 {
		return 1
	}

	cv, ok := cache[target]
	if ok {
		return cv
	}

	rv := 0
	for w := range words {
		if strings.HasPrefix(target, w) {
			rv += possible(words, target[len(w):])
		}
	}

	cache[target] = rv
	return rv
}

func main() {
	words := make(map[string]bool)
	var desired []string
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		bits := strings.Split(scanner.Text(), ", ")
		if len(bits) == 1 {
			if bits[0] != "" {
				desired = append(desired, bits[0])
			}
		} else {
			for _, bit := range bits {
				words[bit] = true
			}
		}
	}

	c := 0
	for _, target := range desired {
		c += possible(words, target)
	}
	fmt.Println(c)
}
