package main

import "fmt"

func subtotal(harga, jumlah int) int {
	// Fungsi perhitungan juga dapat dipakai tanpa mencetak laporan.
	return harga * jumlah
}

func main() {
	total := subtotal(5000, 3)
	fmt.Printf("total=%d\n", total)
}
