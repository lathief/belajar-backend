package main

import "fmt"

func main() {
	daftar := []int{0, 1, 4, 8}
	valid := 0
	for _, n := range daftar {
		// Kedua batas termasuk jumlah pesanan yang diterima.
		if n >= 1 && n <= 5 {
			valid++
		}
	}
	fmt.Printf("valid=%d\n", valid)
}
