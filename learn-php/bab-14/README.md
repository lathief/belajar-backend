# Bab 14 — Ekspor audit dengan paginasi cursor

Paginasi berbasis cursor melanjutkan pembacaan setelah ID terakhir, tidak bergantung pada offset yang bergeser saat ada data baru.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
halaman=E1,E2 cursor=E2
halaman=E3,E4 cursor=E4
halaman=E5 cursor=E5
```

## Telusuri program

Setelah halaman pertama, data baru E5 ditambahkan. Cursor E2 tetap membuat halaman selanjutnya mulai dari E3 dan setiap ID muncul tepat sekali.

## Latihan mandiri

Ubah ukuran halaman menjadi 3 dan periksa urutan semua ID.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-13/README.md) · [Bab berikutnya](../bab-15/README.md)
