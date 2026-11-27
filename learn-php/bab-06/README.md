# Bab 06 — Class, property, dan method

Class menggabungkan data dan perilaku. Property `private` membatasi perubahan stok agar hanya dilakukan lewat method yang memeriksa aturan.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
diterima stok=3
stok_kurang stok=3
```

## Telusuri program

`Gudang::pesan` mengurangi stok hanya bila jumlah positif dan stok cukup. Pemanggil membaca stok lewat method `stok()`.

## Latihan mandiri

Coba pesan dengan jumlah `0`; method harus menolak tanpa mengubah stok.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-05/README.md) · [Bab berikutnya](../bab-07/README.md)
