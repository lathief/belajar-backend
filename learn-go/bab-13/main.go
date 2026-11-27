package main

import "fmt"

type peristiwa struct{ nomor, delta int }
type ringkasan struct{ posisi, stok int }

func (r *ringkasan) terapkan(e peristiwa) string {
	switch {
	case e.nomor <= r.posisi:
		return "duplikat"
	case e.nomor != r.posisi+1:
		return "celah"
	}
	// Majukan posisi hanya bersama perubahan stok.
	r.stok += e.delta
	r.posisi = e.nomor
	return "diterapkan"
}

func main() {
	var r ringkasan
	for _, e := range []peristiwa{{1, 5}, {3, -2}, {2, -1}, {2, -1}, {3, -2}} {
		hasil := r.terapkan(e)
		fmt.Printf("%d %s posisi=%d stok=%d\n", e.nomor, hasil, r.posisi, r.stok)
	}
}
