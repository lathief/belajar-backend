package main

import "fmt"

func kirim(respons []int, deadline int) {
	waktu, jeda := 0, 1
	for i, kode := range respons {
		fmt.Printf("t=%d percobaan=%d kode=%d\n", waktu, i+1, kode)
		if kode == 200 {
			fmt.Println("hasil=sukses")
			return
		}
		if kode != 503 {
			fmt.Println("hasil=permanen")
			return
		}
		if i+1 == len(respons) {
			fmt.Println("hasil=habis")
			return
		}
		// Batas eksklusif: percobaan baru harus mulai sebelum deadline.
		if waktu+jeda >= deadline {
			fmt.Println("hasil=deadline")
			return
		}
		waktu += jeda
		jeda *= 2
	}
}

func main() {
	kirim([]int{503, 503, 200}, 4)
	kirim([]int{503, 503, 200}, 3)
	kirim([]int{400, 200}, 4)
}
