package main

import "fmt"

type Ongkir interface {
	Hitung(subtotal int) int
}

type TarifTetap struct{ Biaya int }
type AmbilSendiri struct{}

func (t TarifTetap) Hitung(_ int) int { return t.Biaya }
func (AmbilSendiri) Hitung(_ int) int { return 0 }

func tagihan(subtotal int, aturan Ongkir) int {
	// Pemanggil memilih aturan, perhitungan memakai kontrak yang sama.
	return subtotal + aturan.Hitung(subtotal)
}

func main() {
	fmt.Printf("kirim=%d\n", tagihan(15000, TarifTetap{Biaya: 2000}))
	fmt.Printf("ambil=%d\n", tagihan(15000, AmbilSendiri{}))
}
