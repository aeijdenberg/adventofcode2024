package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	type coord struct{ x, y int }

	obstacles := make(map[coord]bool)

	var gx, gy, gd, w, h int
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); h++ {
		w = 0
		for _, b := range scanner.Bytes() {
			switch b {
			case '#':
				obstacles[coord{w, h}] = true
			case '^':
				gx, gy, gd = w, h, 0
			case '.':
			default:
				panic(b)
			}
			w++
		}
	}

	dirs := []coord{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}

	visited := make(map[coord]bool)
	for gx >= 0 && gx < w && gy >= 0 && gy < h {
		visited[coord{gx, gy}] = true
		for obstacles[coord{gx + dirs[gd].x, gy + dirs[gd].y}] {
			gd = (gd + 1) % len(dirs)
		}
		gx, gy = gx+dirs[gd].x, gy+dirs[gd].y
	}

	fmt.Println(len(visited))
}
