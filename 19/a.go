package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func possible(words map[string]bool, target string) bool {
	if len(target) == 0 {
		return true
	}
	for w := range words {
		if strings.HasPrefix(target, w) {
			rv := possible(words, target[len(w):])
			if rv {
				return true
			}
		}
	}
	return false
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
		if possible(words, target) {
			c++
		}
	}
	fmt.Println(c)
}
