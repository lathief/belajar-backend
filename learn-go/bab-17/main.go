package main

import "fmt"

type tugas struct {
	generasi, kedaluwarsa int
	selesai               bool
}

func (t *tugas) klaim(kini int) (int, bool) {
	if t.selesai || kini < t.kedaluwarsa {
		return 0, false
	}
	t.generasi++
	t.kedaluwarsa = kini + 5
	return t.generasi, true
}

func (t *tugas) simpan(token, kini int) string {
	// Periksa waktu juga ketika pengambil alih belum datang.
	if token == 0 || t.selesai || token != t.generasi || kini >= t.kedaluwarsa {
		return "ditolak"
	}
	t.selesai = true
	return "tersimpan"
}

func main() {
	var t tugas
	a, _ := t.klaim(0)
	_, dapat := t.klaim(4)
	fmt.Printf("klaim_dini=%t\n", dapat)
	fmt.Println("batas=" + t.simpan(a, 5))
	b, _ := t.klaim(5)
	fmt.Println("lama=" + t.simpan(a, 6))
	fmt.Println("baru=" + t.simpan(b, 6))
	fmt.Println("ulang=" + t.simpan(b, 7))
}
