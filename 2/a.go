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
	scanner := bufio.NewScanner(os.Stdin)
	var all [][]int
	for scanner.Scan() {
		bits := strings.Split(scanner.Text(), " ")
		nums := make([]int, len(bits))
		for i, b := range bits {
			nums[i] = MustInt(b)
		}
		all = append(all, nums)
	}

	safe := 0
	for _, row := range all {
		var countDesc, countAsc, wildRange int
		for i := 0; i < len(row)-1; i++ {
			d := row[i+1] - row[i]
			if d > 0 { // we are ascending
				countAsc++
				if d > 3 {
					wildRange++
				}
			} else if d < 0 {
				countDesc++
				if d < -3 {
					wildRange++
				}
			}
		}
		if wildRange == 0 {
			if countDesc == len(row)-1 || countAsc == len(row)-1 {
				safe++
			}
		}
	}

	fmt.Println(safe)
}
