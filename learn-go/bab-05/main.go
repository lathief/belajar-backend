package main

import "fmt"

func main() {
	kode := []string{"PS-1", "PS-2"}
	kode = append(kode, "PS-1")
	stok := map[string]int{"PS-1": 4, "PS-2": 0}
	// Boolean hasil lookup membedakan stok nol dari kode yang belum tercatat.
	sisa, ada := stok["PS-2"]
	fmt.Printf("entri=%d\n", len(kode))
	fmt.Printf("stok=%d ada=%t\n", sisa, ada)
}
