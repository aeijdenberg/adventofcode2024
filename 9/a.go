package main

import (
	"fmt"
	"io"
	"os"
)

type chunk struct {
	Prev *chunk
	Next *chunk

	BlockStart int

	File   bool
	ID     int
	Length int
}

func PrintIt(first *chunk) {
	for first != nil {
		for i := 0; i < first.Length; i++ {
			if first.File {
				fmt.Printf("%d", first.ID)
			} else {
				fmt.Printf(".")
			}
		}
		first = first.Next
	}
	fmt.Printf("\n")
}

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	first := &chunk{}
	last := first

	nextFile := true
	nextID := 0
	blockStart := 0
	for _, c := range b {
		last.Next = &chunk{
			BlockStart: blockStart,
			Prev:       last,
			Length:     int(c - '0'),
			File:       nextFile,
		}
		blockStart += last.Next.Length
		last = last.Next
		if nextFile {
			last.ID = nextID
			nextID++
		}
		nextFile = !nextFile
	}
	last.Next = &chunk{Prev: last, BlockStart: blockStart}
	last = last.Next

	// ok, now go from both ends
	left, right := first, last
	for {
		// PrintIt(first)

		for left != nil && (left.Length == 0 || left.File) {
			left = left.Next
		}

		for right != nil && (right.Length == 0 || !right.File) {
			right = right.Prev
		}

		if left == nil || right == nil {
			break
		}
		if left.BlockStart >= right.BlockStart {
			break
		}

		amtToConsume := min(left.Length, right.Length)

		// need to split left
		newLeft := &chunk{
			BlockStart: left.BlockStart,
			Prev:       left.Prev,
			Next:       left,

			File:   true,
			ID:     right.ID,
			Length: amtToConsume,
		}
		left.Prev.Next = newLeft
		left.Prev = newLeft
		left.BlockStart += amtToConsume
		left.Length -= amtToConsume

		// need to split right too!
		newRight := &chunk{
			BlockStart: right.BlockStart + amtToConsume,
			Prev:       right,
			Next:       right.Next,
			Length:     amtToConsume,
		}
		right.Next.Prev = newRight
		right.Next = newRight
		right.Length -= amtToConsume
	}

	PrintIt(first)

	s := 0
	for cur := first; cur != last; cur = cur.Next {
		if cur.File {
			for i := 0; i < cur.Length; i++ {
				s += (cur.BlockStart + i) * cur.ID
			}
		}
	}

	fmt.Println(s)
}
