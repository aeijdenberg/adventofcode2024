package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type C struct {
	X, Y int
}

var dirs = []C{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func shortestPath(empties map[C]bool, s, e, e1 C, target int) int {
	bestToE := math.MaxInt

	// dijstra this
	source := s
	dist := make(map[C]int)
	dist[source] = 0
	q := &PQ{}
	q.AddWithPriority(source, 0)
	for !q.Empty() {
		u := q.Pop()

		for _, dir := range dirs {
			v := C{u.X + dir.X, u.Y + dir.Y}
			if !empties[v] && v != e1 {
				continue
			}

			alt := dist[u] + 1
			if alt > target {
				continue
			}
			// else we are good
			distToV, found := dist[v]
			if !found || alt < distToV {
				dist[v] = alt
				q.AddWithPriority(v, alt)
				if v == e {
					bestToE = min(bestToE, alt)
				}
			}
		}
	}
	return bestToE
}

func (c C) InFrame(w, h int) bool {
	if c.X < 1 || c.X >= (w-1) {
		return false
	}
	if c.Y < 1 || c.Y >= (h-1) {
		return false
	}
	return true
}

func main() {
	empties := make(map[C]bool)
	walls := make(map[C]bool)
	var e, s C
	var w, h int
	for y, scanner := 0, bufio.NewScanner(os.Stdin); scanner.Scan(); y++ {
		w = 0
		for x, ch := range scanner.Text() {
			c := C{x, y}
			switch ch {
			case '.':
				empties[c] = true
			case '#':
				walls[c] = true
			case 'S':
				s = c
				empties[c] = true
			case 'E':
				e = c
				empties[c] = true
			}
			w++
		}
		h++
	}

	bestWithNoCheat := shortestPath(empties, s, e, e, math.MaxInt)

	answer := 0
	for c1 := range walls {
		if !c1.InFrame(w, h) { // || !c2.InFrame(w, h) {
			continue
		}
		bestWithCheat := shortestPath(empties, s, e, c1, bestWithNoCheat-100)
		if bestWithCheat != math.MaxInt {
			answer++
		}
	}

	fmt.Println(answer)
}

type PQ struct {
	q *pnode
}

type pnode struct {
	v    C
	p    int
	next *pnode
}

func (pq *PQ) AddWithPriority(v C, d int) {
	// fmt.Println("add", v, d)
	n := &pnode{
		v: v,
		p: d,
	}

	if pq.q == nil {
		pq.q = n
		return
	}

	var last *pnode
	cur := pq.q
	for cur != nil && cur.p > d {
		last = cur
		cur = cur.next
	}

	n.next = cur
	if last == nil {
		pq.q = n
		return
	}
	last.next = n
}

func (pq *PQ) Pop() C {
	rv := pq.q
	pq.q = rv.next
	return rv.v
}

func (pq *PQ) Empty() bool {
	return pq.q == nil
}
