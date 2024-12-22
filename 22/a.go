package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MustUint(s string) uint64 {
	rv, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return rv
}

func Next(secret uint64) uint64 {
	secret = (secret ^ (secret << 6)) & 0xffffff
	secret = (secret ^ (secret >> 5)) & 0xffffff
	secret = (secret ^ (secret << 11)) & 0xffffff
	return secret
}

func main() {
	s := uint64(0)
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		n := MustUint(scanner.Text())
		for i := 0; i < 2000; i++ {
			n = Next(n)
		}
		s += n
	}
	fmt.Println(s)
}
