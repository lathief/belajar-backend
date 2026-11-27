# Bab 05 — Slice dan map

Slice menyimpan urutan nilai dan bisa tumbuh. Simpan kembali hasil `append` karena operasi itu dapat mengalokasikan array baru. Map menyimpan pasangan kunci dan nilai. Bentuk `nilai, ada := m[kunci]` membedakan data yang tersimpan dari kunci yang hilang.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-05`, jalankan perintah berikut.

```sh
go run main.go
```

Keluaran program yang diharapkan

```text
entri=3
stok=0 ada=true
```

## Telusuri program

Setelah `append`, daftar memiliki tiga entri. Kunci `PS-2` tersimpan dengan stok 0 sehingga `ada` bernilai `true`. Saat pencarian diganti menjadi `PS-9`, nilai stok yang dibaca tetap 0 tetapi `ada` menjadi `false`.

## Latihan mandiri

Tambahkan satu kode ke daftar dan verifikasi panjangnya 4. Cari `PS-9`, lalu gunakan `if !ada` untuk mencetak `kode tidak ditemukan`. Pastikan `PS-2` tetap dianggap ada.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-04/README.md) · [Bab berikutnya](../bab-06/README.md)
