package main

import "fmt"

type produsen struct {
	stok              int
	dipesan, tertunda bool
}

func (p *produsen) pesan() {
	if !p.dipesan {
		// Satu commit model untuk pesanan A dan entri outbox A.
		*p = produsen{stok: p.stok - 1, dipesan: true, tertunda: true}
	}
}

type konsumen struct {
	inbox map[string]bool
	paket int
}

func (k *konsumen) terima(id string) {
	if !k.inbox[id] {
		// Dalam penyimpanan nyata, inbox dan paket berada dalam satu transaksi.
		k.inbox[id] = true
		k.paket++
	}
}

func simulasi(responsHilang, relayPutus, ackHilang bool) (int, int, int) {
	p := produsen{stok: 5}
	p.pesan()
	if responsHilang {
		p.pesan()
	}
	var antrean []string
	kirim := 0
	for p.tertunda {
		antrean = append(antrean, "A")
		kirim++
		if relayPutus {
			relayPutus = false
			continue
		}
		p.tertunda = false
	}
	k := konsumen{inbox: make(map[string]bool)}
	for len(antrean) > 0 {
		k.terima(antrean[0])
		if ackHilang {
			ackHilang = false
			continue
		}
		antrean = antrean[1:]
	}
	return p.stok, kirim, k.paket
}

func main() {
	lulus := 0
	for _, respons := range []bool{false, true} {
		for _, relay := range []bool{false, true} {
			for _, ack := range []bool{false, true} {
				stok, kirim, paket := simulasi(respons, relay, ack)
				target := 1
				if relay {
					target = 2
				}
				if stok != 4 || kirim != target || paket != 1 {
					panic("invarian gagal")
				}
				lulus++
				if respons && relay && ack {
					fmt.Printf("stok=%d publikasi=%d paket=%d\n", stok, kirim, paket)
				}
			}
		}
	}
	fmt.Printf("skenario_lulus=%d\n", lulus)
}
