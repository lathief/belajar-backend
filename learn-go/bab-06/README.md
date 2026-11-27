# Bab 06 — Struct, method, dan pointer

`struct` mengelompokkan field yang berkaitan. Parameter sebelum nama method disebut receiver. Receiver `*Persediaan` memberi method akses ke objek milik pemanggil, sehingga perubahan stok terlihat sesudah pemanggilan.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-06`, jalankan perintah berikut.

```sh
go run main.go
```

Keluaran program yang diharapkan

```text
diambil=true sisa=1
```

## Telusuri program

Pada `stok.Ambil(2)`, Go dapat mengambil alamat variabel `stok` untuk receiver pointer. Method mengurangi stok dari 3 menjadi 1. Jika penerimanya ditulis `p Persediaan`, pengurangan hanya terjadi pada salinan.

## Latihan mandiri

Setelah pengambilan pertama, panggil `stok.Ambil(2)` sekali lagi. Hasilnya harus `false` dan stok tetap 1. Uji juga jumlah 0 untuk memeriksa batas bawah.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-05/README.md) · [Bab berikutnya](../bab-07/README.md)
