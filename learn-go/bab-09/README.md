# Bab 09 — Proyek mini rekapan pesanan

Buat alat terminal yang menerima satu jumlah pesanan per argumen. Terima bilangan bulat 1–5, jumlahkan unit yang sah, dan hitung berapa entri yang ditolak. `os.Args[1:]` melewati nama program. Pisahkan pengolahan data di `rekap` agar aturan bisa dipanggil tanpa terminal.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-09`, jalankan perintah berikut.

```sh
go run main.go 2 0 x 3
```

Keluaran program yang diharapkan

```text
unit=5 ditolak=2
```

## Telusuri program

Untuk `2 0 x 3`, angka 2 menambah unit, 0 gagal pada rentang, `x` gagal saat parsing, dan 3 menambah unit lagi. Jika daftar kosong, loop dilewati dan kedua penghitung tetap nol.

## Latihan mandiri

Tambahkan hasil ketiga untuk menghitung pesanan valid. Uji `2 0 x 3` (2 pesanan valid), tanpa argumen (semua nol), dan `1 5 6` (2 pesanan valid, 6 unit, 1 ditolak).

[Roadmap](../README.md) · [Bab sebelumnya](../bab-08/README.md) · [Bab berikutnya](../bab-10/README.md)
