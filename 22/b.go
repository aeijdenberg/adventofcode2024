package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MustInt(s string) int64 {
	rv, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return rv
}

func Next(secret int64) int64 {
	secret = (secret ^ (secret << 6)) & 0xffffff // does nothing to right most digit
	secret = (secret ^ (secret >> 5)) & 0xffffff
	secret = (secret ^ (secret << 11)) & 0xffffff
	return secret
}

type Buyer map[uint32]int64

func NewBuyer(n int64) Buyer {
	rv := make(Buyer)
	lastMod10 := n % 10
	var deltaKey uint32
	for i := 0; i < 2000; i++ {
		n = Next(n)
		nextMod10 := n % 10
		deltaKey = (deltaKey << 8) | uint32(10+nextMod10-lastMod10)
		lastMod10 = nextMod10

		_, found := rv[deltaKey]
		if (!found) && (i >= 3) {
			rv[deltaKey] = nextMod10
		}
	}
	return rv
}

func ToKey(d1, d2, d3, d4 int) uint32 {
	return uint32(((d1 + 10) << 24) | ((d2 + 10) << 16) | ((d3 + 10) << 8) | (d4 + 10))
}

func main() {
	var all []Buyer
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		all = append(all, NewBuyer(MustInt(scanner.Text())))
	}

	allKeys := make(map[uint32]bool)
	for _, b := range all {
		for k := range b {
			allKeys[k] = true
		}
	}

	bestVal := int64(-1)
	for k := range allKeys {
		a := int64(0)
		for _, b := range all {
			a += b[k]
		}
		bestVal = max(bestVal, a)
	}

	fmt.Println(bestVal)
}
