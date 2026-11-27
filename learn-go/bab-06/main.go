package main

import "fmt"

type Persediaan struct {
	Sisa int
}

func (p *Persediaan) Ambil(jumlah int) bool {
	// Periksa seluruh syarat sebelum mengubah stok.
	if jumlah <= 0 || jumlah > p.Sisa {
		return false
	}
	p.Sisa -= jumlah
	return true
}

func main() {
	stok := Persediaan{Sisa: 3}
	ok := stok.Ambil(2)
	fmt.Printf("diambil=%t sisa=%d\n", ok, stok.Sisa)
}
