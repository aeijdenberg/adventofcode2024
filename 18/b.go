package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func MustInt(s string) int {
	rv, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return rv
}

func feasible(numBytes, w, h int, corrupts map[C]int) bool {
	// dijstra this
	source := C{0, 0}
	dist := make(map[C]int)
	dist[source] = 0
	q := &PQ{}
	q.AddWithPriority(source, 0)
	for !q.Empty() {
		u := q.Pop()

		for _, dir := range dirs {
			v := C{u.X + dir.X, u.Y + dir.Y}

			// is it in range?
			if v.X < 0 || v.X > w || v.Y < 0 || v.Y > h {
				continue
			}

			alt := dist[u] + 1

			// is it corrupt?
			corruptTime, found := corrupts[v]
			if found && corruptTime < numBytes { // TODO, if off by one, try < u.T or <= u.T + 1
				continue
			}
			// else we are good
			distToV, found := dist[v]
			if !found || alt < distToV {
				dist[v] = alt
				q.AddWithPriority(v, alt)
				if (v == C{w, h}) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	corrupts := make(map[C]int)
	var corruptsByIndex []C
	for t, scanner := 0, bufio.NewScanner(os.Stdin); scanner.Scan(); t++ {
		bits := strings.Split(scanner.Text(), ",")
		corrupts[C{MustInt(bits[0]), MustInt(bits[1])}] = t
		corruptsByIndex = append(corruptsByIndex, C{MustInt(bits[0]), MustInt(bits[1])})
	}

	w, h := 0, 0
	for c := range corrupts {
		w = max(w, c.X)
		h = max(h, c.Y)
	}

	firstNotFeasible := sort.Search(len(corruptsByIndex), func(i int) bool {
		return !feasible(i, w, h, corrupts)
	})

	fmt.Printf("%d,%d\n", corruptsByIndex[firstNotFeasible-1].X, corruptsByIndex[firstNotFeasible-1].Y)
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
