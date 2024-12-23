package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type C struct {
	A, B string
}

func main() {
	vs := make(map[string]bool)
	es := make(map[C]bool) // alphabetical

	groups := make(map[string]bool)
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		bits := strings.Split(scanner.Text(), "-")
		slices.Sort(bits)
		es[C{bits[0], bits[1]}] = true
		groups[strings.Join(bits, ",")] = true
		vs[bits[0]] = true
		vs[bits[1]] = true
	}

	done := false
	var lastBest string
	for groupSizeTarget := 2; !done; groupSizeTarget++ {
		done = true
		gstsl := (groupSizeTarget * 3) - 1
		for g := range groups {
			if len(g) == gstsl {
				soFar := make(map[string]bool)
				for _, b := range strings.Split(g, ",") {
					soFar[b] = true
				}
				// ok, now try an additional candidate
				for v := range vs {
					if !soFar[v] {
						// now check that *each* connection exists
						allMatch := true
						for u := range soFar {
							potEdge := []string{v, u}
							slices.Sort(potEdge)
							if !es[C{potEdge[0], potEdge[1]}] {
								allMatch = false
								break
							}
						}
						// TODO, save not to try again?
						if allMatch {
							newGroup := []string{v}
							for u := range soFar {
								newGroup = append(newGroup, u)
							}
							slices.Sort(newGroup)
							lastBest = strings.Join(newGroup, ",")
							// fmt.Println("found", lastBest, len(newGroup))
							groups[lastBest] = true
							done = false
						}
					}
				}
			}
		}
	}
	fmt.Println(lastBest)
}
