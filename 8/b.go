package main

import (
	"bufio"
	"fmt"
	"os"
)

type C struct {
	X, Y int
}

func main() {
	var w, h int
	freqs := make(map[byte][]C)
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); h++ {
		w = 0
		for _, ch := range scanner.Bytes() {
			if ch != '.' {
				freqs[ch] = append(freqs[ch], C{w, h})
			}
			w++
		}
	}

	antinodes := make(map[C]bool)
	for _, cs := range freqs {
		for i := 0; i < len(cs); i++ {
			for j := i + 1; j < len(cs); j++ {
				dx := cs[j].X - cs[i].X
				dy := cs[j].Y - cs[i].Y
				for done, factor := false, 0; !done; factor++ {
					done = true
					for _, nn := range []C{
						{cs[i].X - factor*dx, cs[i].Y - factor*dy},
						{cs[j].X + factor*dx, cs[j].Y + factor*dy},
					} {
						if nn.X >= 0 && nn.X < w && nn.Y >= 0 && nn.Y < h {
							antinodes[nn] = true
							done = false
						}
					}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}
