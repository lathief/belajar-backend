# Bab 02 — Variabel, tipe, dan operator

Gunakan `const` untuk nilai tetap dan `var` untuk variabel. Di dalam fungsi, `:=` mendeklarasikan variabel dengan tipe yang disimpulkan dari nilai awal. Harga contoh disimpan sebagai bilangan bulat dalam rupiah.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-02`, jalankan perintah berikut.

```sh
go run main.go
```

Keluaran program yang diharapkan

```text
subtotal=15000
```

## Telusuri program

Nilai awal `jumlah` adalah 2. Setelah `jumlah++`, perkalian memakai nilai 3. Pada `Printf`, `%d` menampilkan bilangan bulat dan `\n` mengakhiri baris.

## Latihan mandiri

Ubah jumlah awal menjadi 3. Hasilnya harus `subtotal=20000`. Lalu coba `var jumlah int` tanpa nilai awal dan jelaskan mengapa hasilnya 5000.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-01/README.md) · [Bab berikutnya](../bab-03/README.md)
