package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct{ x, y int }

func isLoop(obstacles map[coord]bool, ox, oy, gx, gy, gd, w, h int) bool {
	obstacles[coord{ox, oy}] = true
	defer delete(obstacles, coord{ox, oy})

	dirs := []coord{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}

	visited := make(map[coord]int)
	for gx >= 0 && gx < w && gy >= 0 && gy < h {
		if visited[coord{gx, gy}] == gd+1 {
			return true
		}
		visited[coord{gx, gy}] = gd + 1
		for obstacles[coord{gx + dirs[gd].x, gy + dirs[gd].y}] {
			gd = (gd + 1) % len(dirs)
		}
		gx, gy = gx+dirs[gd].x, gy+dirs[gd].y
	}
	return false
}

func main() {

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

	c := 0
	for ox := 0; ox < w; ox++ {
		for oy := 0; oy < h; oy++ {
			if obstacles[coord{ox, oy}] {
				continue
			}
			if gx == ox && gy == oy {
				continue
			}
			if isLoop(obstacles, ox, oy, gx, gy, gd, w, h) {
				c++
			}
		}
	}

	fmt.Println(c)
}
