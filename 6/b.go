package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct{ x, y int }

var dirs = []coord{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func getVisits(obstacles map[coord]bool, gx, gy, gd, w, h int) map[coord]bool {
	visited := make(map[coord]bool)
	for gx >= 0 && gx < w && gy >= 0 && gy < h {
		visited[coord{gx, gy}] = true
		for obstacles[coord{gx + dirs[gd].x, gy + dirs[gd].y}] {
			gd = (gd + 1) % len(dirs)
		}
		gx, gy = gx+dirs[gd].x, gy+dirs[gd].y
	}
	return visited
}

func isLoop(obstacles map[coord]bool, ox, oy, gx, gy, gd, w, h int, results chan int) {
	visited := make(map[coord]int)
	for gx >= 0 && gx < w && gy >= 0 && gy < h {
		if visited[coord{gx, gy}] == gd+1 {
			results <- 1
			return
		}
		visited[coord{gx, gy}] = gd + 1
		for obstacles[coord{gx + dirs[gd].x, gy + dirs[gd].y}] || (gx+dirs[gd].x == ox && gy+dirs[gd].y == oy) {
			gd = (gd + 1) % len(dirs)
		}
		gx, gy = gx+dirs[gd].x, gy+dirs[gd].y
	}
	results <- 0
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

	results := make(chan int, 32)
	remaining := 0
	for oc := range getVisits(obstacles, gx, gy, gd, w, h) {
		if gx == oc.x && gy == oc.y {
			continue
		}
		remaining++
		go isLoop(obstacles, oc.x, oc.y, gx, gy, gd, w, h, results)
	}

	c := 0
	for ; remaining != 0; remaining-- {
		c += <-results
	}

	fmt.Println(c)
}
