# Bab 05 — Array berindeks dan asosiatif

Array PHP dapat berindeks angka atau memakai key bernama. Operator `??` memberi nilai cadangan ketika key tidak ditemukan.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
PS-1 stok=0
KODE-HILANG stok=tidak_tercatat
```

## Telusuri program

Stok nol adalah data yang ada, bukan key yang hilang. Ekspresi `$stok[$kode] ?? ...` mempertahankan perbedaan itu.

## Latihan mandiri

Tambahkan key `PS-2` dengan stok `4` dan cetak nilainya.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-04/README.md) · [Bab berikutnya](../bab-06/README.md)
