package main

import "fmt"

type receipt struct {
	jumlah int
	hasil  string
}

type gudang struct {
	stok    int
	riwayat map[string]receipt
}

func (g *gudang) pesan(id string, jumlah int) string {
	if jumlah < 1 || jumlah > 5 {
		return "jumlah_invalid"
	}
	// Baca keputusan pertama sebelum menilai stok yang mungkin sudah berubah.
	if awal, ada := g.riwayat[id]; ada {
		if awal.jumlah != jumlah {
			return "konflik"
		}
		return awal.hasil
	}

	hasil := "stok_kurang"
	if jumlah <= g.stok {
		g.stok -= jumlah
		hasil = "diterima"
	}
	g.riwayat[id] = receipt{jumlah: jumlah, hasil: hasil}
	return hasil
}

func main() {
	g := gudang{stok: 5, riwayat: make(map[string]receipt)}
	urutan := []struct {
		id     string
		jumlah int
	}{
		{"A", 3}, {"B", 3}, {"C", 0}, {"C", 2},
		{"A", 3}, {"A", 1}, {"B", 3},
	}
	for _, r := range urutan {
		hasil := g.pesan(r.id, r.jumlah)
		fmt.Printf("%s %d %s stok=%d\n", r.id, r.jumlah, hasil, g.stok)
	}
}
