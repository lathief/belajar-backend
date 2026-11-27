package main

import (
	"fmt"
	"os"
	"strconv"
)

func rekap(daftar []string) (int, int) {
	unit, ditolak := 0, 0
	for _, teks := range daftar {
		n, err := strconv.Atoi(teks)
		if err != nil || n < 1 || n > 5 {
			// Input buruk hanya menolak satu entri dalam batch.
			ditolak++
			continue
		}
		unit += n
	}
	return unit, ditolak
}

func main() {
	unit, ditolak := rekap(os.Args[1:])
	fmt.Printf("unit=%d ditolak=%d\n", unit, ditolak)
}
