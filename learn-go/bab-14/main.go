package main

import (
	"fmt"
	"strconv"
	"strings"
)

func halaman(log []int, setelah, batas, ukuran int) []int {
	var hasil []int
	for _, id := range log {
		if id > batas || len(hasil) == ukuran {
			break
		}
		if id > setelah {
			hasil = append(hasil, id)
		}
	}
	return hasil
}

func tampil(ids []int) string {
	var teks []string
	for _, id := range ids {
		teks = append(teks, strconv.Itoa(id))
	}
	return strings.Join(teks, ",")
}

func main() {
	log := []int{101, 104, 107, 112}
	batas := log[len(log)-1]
	cursor := 0
	for {
		ids := halaman(log, cursor, batas, 2)
		if len(ids) == 0 {
			break
		}
		fmt.Printf("halaman=%s\n", tampil(ids))
		if cursor == 0 {
			log = append(log, 115)
		} // Datang setelah halaman pertama.
		cursor = ids[len(ids)-1]
	}
	batasBaru := log[len(log)-1]
	fmt.Printf("susulan=%s\n", tampil(halaman(log, cursor, batasBaru, 2)))
}
