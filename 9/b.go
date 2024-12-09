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

	for right := last; right != nil; right = right.Prev {
		if !right.File {
			continue
		}
		if right.Length == 0 {
			continue
		}

		// else we have a file, now search from the left
		var spaceStart *chunk
		var freeSpace int
		for left := first; left != nil && left.BlockStart < right.BlockStart; left = left.Next {
			if left.File {
				spaceStart = nil
				freeSpace = 0
			} else if left.Length != 0 {
				if spaceStart == nil {
					spaceStart = left
					freeSpace = 0
				}
				freeSpace += left.Length
				if freeSpace >= right.Length {
					// MOVE IT
					newFile := &chunk{
						BlockStart: spaceStart.BlockStart,
						Prev:       spaceStart.Prev,
						File:       true,
						ID:         right.ID,
						Length:     right.Length,
					}
					newEmpty := &chunk{
						BlockStart: spaceStart.BlockStart + right.Length,
						Prev:       newFile,
						Length:     freeSpace - right.Length,
						Next:       left.Next,
					}
					newFile.Next = newEmpty

					spaceStart.Prev.Next = newFile
					left.Next.Prev = newEmpty

					right.File = false
					break
				}
			}
		}
	}

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
