# Bab 09 — Proyek mini rekapan pesanan

`process.argv.slice(2)` mengambil argumen sesudah executable Node dan lokasi modul. Pakai kembali validator bab 7. Program memproses tiap argumen secara independen dan tetap membuat laporan ketika salah satu pesanan ditolak.

## Baca dan jalankan

Buka [main.mjs](main.mjs) dan [validator dari bab 07](../bab-07/validasi.mjs). Dari direktori `learn-javascript/bab-09`, jalankan perintah berikut.

```sh
node main.mjs 2 0 x 3
```

Keluaran program yang diharapkan

```text
unit=5 ditolak=2
```

## Telusuri program

Dua elemen awal `process.argv` dipakai oleh Node. Dari empat argumen pengguna, `2` dan `3` menambah unit, sedangkan `0` dan `x` masing-masing menambah satu penolakan.

## Latihan mandiri

Tambahkan hitungan pesanan valid untuk `2 0 x 3` (2 pesanan). Jalankan tanpa argumen (kedua angka nol) dan dengan `1 5 6` (6 unit, 1 ditolak).

[Roadmap](../README.md) · [Bab sebelumnya](../bab-08/README.md) · [Bab berikutnya](../bab-10/README.md)
