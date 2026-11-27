# Bab 04 — Fungsi dan nilai balik

Setelah nama fungsi, tulis parameter di dalam kurung dan tipe hasil sesudahnya. Parameter berurutan dengan tipe sama dapat ditulis `harga, jumlah int`. `return` mengirim hasil ke pemanggil.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-04`, jalankan perintah berikut.

```sh
go run main.go
```

Keluaran program yang diharapkan

```text
total=15000
```

## Telusuri program

Pemanggilan `subtotal(5000, 3)` mengikat 5000 ke `harga` dan 3 ke `jumlah`. Variabel parameter hanya tersedia di fungsi itu. `main` menerima hasilnya dan memilih cara menampilkannya.

## Latihan mandiri

Tambahkan parameter `diskon` dalam rupiah. Untuk harga 5000, jumlah 3, dan diskon 2000, total harus 13000. Jika diskon melebihi subtotal, kembalikan 0.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-03/README.md) · [Bab berikutnya](../bab-05/README.md)
