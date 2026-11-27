# Bab 12 — Konflik penulisan dan nomor versi

Dua klien dapat membaca versi yang sama lalu mengirim perubahan. Pemeriksaan versi optimistis menerima hanya perubahan yang didasarkan pada versi terkini.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
tersimpan versi=2
konflik versi_terkini=2
stok=8 versi=2
```

## Telusuri program

Penulisan pertama memakai versi 1 dan menaikkan versi menjadi 2. Formulir kedua masih memakai versi 1, jadi perubahan itu ditolak tanpa menimpa stok.

## Latihan mandiri

Ulangi penulisan kedua dengan versi 2. Periksa versi baru dan stok hasilnya.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-11/README.md) · [Bab berikutnya](../bab-13/README.md)
