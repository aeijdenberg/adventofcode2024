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

type V struct {
	C C
	D int // 0 = E, 1 = S, 2 = W, 3 = N
}

type N struct {
	V V
	C int // cost
}

func addAllPrevs(allThings map[C]bool, equalBestPrevs map[V]map[V]bool, v V) {
	allThings[v.C] = true
	for u := range equalBestPrevs[v] {
		addAllPrevs(allThings, equalBestPrevs, u)
	}
}

func main() {
	empties := make(map[C]bool)
	var end, start C
	for y, scanner := 0, bufio.NewScanner(os.Stdin); scanner.Scan(); y++ {
		for x, ch := range scanner.Text() {
			switch ch {
			case 'E':
				empties[C{x, y}] = true
				end = C{x, y}
			case 'S':
				start = C{x, y}
			case '.':
				empties[C{x, y}] = true
			}
		}
	}

	bestToE := math.MaxInt

	// dijstra this
	source := V{C: start, D: 0}
	equalBestPrevs := make(map[V]map[V]bool)
	dist := make(map[V]int)
	dist[source] = 0
	q := &PQ{}
	q.AddWithPriority(source, 0)
	neighbours := make([]N, 0, 3)
	for !q.Empty() {
		u := q.Pop()
		// neighbours are LEFT, RIGHT, next spot same dir
		neighbours = append(neighbours[:0],
			N{
				V: V{C: u.C, D: (u.D + 1) % 4},
				C: 1000,
			},
			N{
				V: V{C: u.C, D: (u.D + 4 - 1) % 4},
				C: 1000,
			})
		// is free in our dir?
		next := C{u.C.X + dirs[u.D].X, u.C.Y + dirs[u.D].Y}
		if empties[next] {
			neighbours = append(neighbours, N{
				V: V{C: next, D: u.D},
				C: 1,
			})
		}

		for _, v := range neighbours {
			alt := dist[u] + v.C
			distToV, found := dist[v.V]

			if !found || alt <= distToV {
				// add to eq best
				if !found || alt < distToV {
					equalBestPrevs[v.V] = map[V]bool{u: true}
				} else { // we are EQUAL
					equalBestPrevs[v.V][u] = true
				}

				if !found || alt < distToV {
					dist[v.V] = alt
					q.AddWithPriority(v.V, alt)
					if v.V.C == end {
						bestToE = min(bestToE, alt)
					}
				}
			}
		}
	}

	allThings := make(map[C]bool)
	for k, v := range dist {
		if k.C == end && v == bestToE {
			addAllPrevs(allThings, equalBestPrevs, k)
		}
	}

	// OK, now we have figured it out, would it run faster second time?
	fmt.Println("First:", bestToE, "Seats:", len(allThings))

}

type PQ struct {
	q *pnode
}

type pnode struct {
	v    V
	p    int
	next *pnode
}

func (pq *PQ) AddWithPriority(v V, d int) {
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

func (pq *PQ) Pop() V {
	rv := pq.q
	pq.q = rv.next
	return rv.v
}

func (pq *PQ) Empty() bool {
	return pq.q == nil
}
