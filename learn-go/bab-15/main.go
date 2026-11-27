package main

import "fmt"

type ember struct {
	token, kapasitas, terakhir int
}

func (e *ember) ambil(detik int) string {
	// Satu token per detik. Jam contoh bergerak monoton.
	e.token = min(e.kapasitas, e.token+detik-e.terakhir)
	e.terakhir = detik
	if e.token == 0 {
		return "tunda"
	}
	e.token--
	return "lolos"
}

func main() {
	e := ember{token: 2, kapasitas: 2}
	for _, t := range []int{0, 0, 0, 1, 6, 6, 6} {
		hasil := e.ambil(t)
		fmt.Printf("t=%d %s token=%d\n", t, hasil, e.token)
	}
}
