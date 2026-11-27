package main

import (
	"errors"
	"fmt"
	"strconv"
)

func bacaJumlah(teks string) (int, error) {
	n, err := strconv.Atoi(teks)
	if err != nil {
		return 0, errors.New("bukan bilangan bulat")
	}
	// Angka yang berhasil dibaca masih harus memenuhi aturan pesanan.
	if n < 1 || n > 5 {
		return 0, errors.New("di luar batas")
	}
	return n, nil
}

func main() {
	for _, teks := range []string{"2", "0", "x"} {
		n, err := bacaJumlah(teks)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("jumlah=%d\n", n)
	}
}
