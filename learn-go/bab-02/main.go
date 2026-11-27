package main

import "fmt"

func main() {
	const harga = 5000
	var jumlah int = 2
	// Satu barang ditambahkan sebelum pesanan ditagih.
	jumlah++
	subtotal := harga * jumlah
	fmt.Printf("subtotal=%d\n", subtotal)
}
