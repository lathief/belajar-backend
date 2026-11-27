# Bab 09 — Proyek mini rekapan pesanan

`$argv` berisi argumen CLI; elemen pertama adalah nama skrip. Program ini memakai kembali validator Bab 07, menghitung unit yang valid, dan tetap melaporkan input yang ditolak.

## Baca dan jalankan

Dari direktori `learn-php/bab-09`:

```sh
php main.php 2 0 x 3
```

Keluaran yang diharapkan:

```text
unit=5 ditolak=2
```

## Telusuri program

Argumen `2` dan `3` diterima. Nilai `0` dan `x` ditolak; masing-masing menambah satu penolakan.

## Latihan mandiri

Jalankan tanpa argumen dan dengan `1 5 6`. Prediksi jumlah unit dan penolakan.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-08/README.md) · [Bab berikutnya](../bab-10/README.md)
