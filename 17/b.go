package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type C struct {
	A, B, C int
	IP      int
	P       []int
	O       []int
}

func (c *C) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		panic("badder")
	}
}

func (c *C) Run() {
	for c.IP < len(c.P) {
		// fmt.Printf("%#v\n", c)
		opcode := c.P[c.IP]
		c.IP++
		operand := c.P[c.IP]
		c.IP++
		switch opcode {
		case 0: // adv
			c.A = c.A >> c.combo(operand)
		case 1: // bxl
			c.B = c.B ^ operand
		case 2: // bst
			c.B = c.combo(operand) & 0x7
		case 3: // jnz
			if c.A != 0 {
				c.IP = operand
			}
		case 4: // bxc
			c.B = c.B ^ c.C
		case 5: // out
			c.O = append(c.O, c.combo(operand)&0x7)
		case 6: // bdv
			c.B = c.A >> c.combo(operand)
		case 7: // cdv
			c.C = c.A >> c.combo(operand)
		default:
			panic("bad")
		}
	}
}

func MustInt(s string) int {
	rv, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return rv
}

func main() {
	c := &C{}
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		s := scanner.Text()
		switch {
		case strings.HasPrefix(s, "Register A:"):
			c.A = MustInt(strings.Split(s, ": ")[1])
		case strings.HasPrefix(s, "Register B:"):
			c.B = MustInt(strings.Split(s, ": ")[1])
		case strings.HasPrefix(s, "Register C:"):
			c.C = MustInt(strings.Split(s, ": ")[1])
		case strings.HasPrefix(s, "Program:"):
			for _, b := range strings.Split(strings.Split(s, ": ")[1], ",") {
				c.P = append(c.P, MustInt(b))
			}
		}
	}
	fmt.Println(c.tryTheSolveHere(len(c.P)-1, 0))
}

func (c *C) tryTheSolveHere(digit, aSolved int) int {
	for i := 0; i < 8; i++ {
		aCand := aSolved | (i << (digit * 3))
		c.A = aCand
		c.B = 0
		c.C = 0
		c.IP = 0
		c.O = nil

		c.Run()

		var s []string
		for _, o := range c.O {
			s = append(s, strconv.Itoa(o))
		}
		if len(c.O) >= digit && c.O[digit] == c.P[digit] {
			if digit == 0 {
				return aCand // we are done!
			} else {
				pRv := c.tryTheSolveHere(digit-1, aCand)
				if pRv != -1 {
					return pRv
				}
			}
		}
	}
	return -1
}
