package main

import "fmt"

type alamat struct {
	kota  string
	versi int
}

func (a *alamat) ubah(versiDibaca int, kota string) string {
	if a.versi != versiDibaca {
		return "konflik"
	}
	// Pemilik data ini memproses satu perintah sampai selesai.
	a.kota = kota
	a.versi++
	return "tersimpan"
}

func main() {
	a := alamat{kota: "Bandung", versi: 1}
	bacaA, bacaB := a.versi, a.versi
	lapor := func(pelaku string, versi int, kota string) {
		hasil := a.ubah(versi, kota)
		fmt.Printf("%s %s kota=%s versi=%d\n", pelaku, hasil, a.kota, a.versi)
	}
	lapor("A", bacaA, "Solo")
	lapor("B", bacaB, "Bogor")
	// B membaca ulang dan pengguna mengonfirmasi alamat tujuan.
	lapor("B", a.versi, "Bogor")
}
