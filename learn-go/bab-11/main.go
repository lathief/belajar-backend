package main

import "fmt"

type status string

const (
	baru    status = "baru"
	lunas   status = "lunas"
	dikirim status = "dikirim"
	batal   status = "batal"
)

func transisi(awal status, aksi string) (status, bool) {
	// Hanya pasangan status dan aksi berikut yang mengubah pesanan.
	switch {
	case awal == baru && aksi == "bayar":
		return lunas, true
	case awal == baru && aksi == "batal":
		return batal, true
	case awal == lunas && aksi == "kirim":
		return dikirim, true
	default:
		return awal, false
	}
}

func main() {
	saatIni := baru
	for _, aksi := range []string{"kirim", "bayar", "batal", "kirim", "bayar"} {
		berikut, ok := transisi(saatIni, aksi)
		fmt.Printf("%s %s -> %s ok=%t\n", saatIni, aksi, berikut, ok)
		saatIni = berikut
	}
}
