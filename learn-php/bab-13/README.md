# Bab 13 — Membangun ulang stok dari log peristiwa

Log peristiwa dapat diputar ulang untuk membangun state. Proyeksi ini menolak ID duplikat, memastikan nomor urut berkelanjutan, dan tidak mengubah state pada kejadian yang invalid.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
stok=7 diterapkan=2
duplikat=E2
urutan_terlewat=4 harapan=3
```

## Telusuri program

Peristiwa E1 dan E2 diterapkan berurutan. E2 yang diulang diabaikan; event urutan 4 ditolak karena urutan 3 belum ada.

## Latihan mandiri

Tambahkan event urutan 3 yang valid, lalu coba kembali event urutan 4.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-12/README.md) · [Bab berikutnya](../bab-14/README.md)
