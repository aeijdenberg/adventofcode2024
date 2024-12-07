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

func checkIt(answer int, nums []int) int {
	if len(nums) == 1 {
		if answer == nums[0] {
			return answer
		} else {
			return 0
		}
	}

	rv := checkIt(answer-nums[len(nums)-1], nums[:len(nums)-1])
	if rv != 0 {
		return answer
	}

	if answer%nums[len(nums)-1] != 0 {
		return 0
	}

	rv = checkIt(answer/nums[len(nums)-1], nums[:len(nums)-1])
	if rv != 0 {
		return answer
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	s := 0
	for scanner.Scan() {
		bits := strings.Split(scanner.Text(), ": ")
		var nums []int
		for _, x := range strings.Split(bits[1], " ") {
			nums = append(nums, MustInt(x))
		}
		s += checkIt(MustInt(bits[0]), nums)
	}
	fmt.Println(s)
}
