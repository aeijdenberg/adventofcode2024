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
			c.A = c.A / (1 << c.combo(operand))
		case 1: // bxl
			c.B = c.B ^ operand
		case 2: // bst
			c.B = c.combo(operand) % 8
		case 3: // jnz
			if c.A != 0 {
				c.IP = operand
			}
		case 4: // bxc
			c.B = c.B ^ c.C
		case 5: // out
			c.O = append(c.O, c.combo(operand)%8)
		case 6: // bdv
			c.B = c.A / (1 << c.combo(operand))
		case 7: // cdv
			c.C = c.A / (1 << c.combo(operand))
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
	c.Run()

	var s []string
	for _, o := range c.O {
		s = append(s, strconv.Itoa(o))
	}
	fmt.Println(strings.Join(s, ","))

}
