# Bab 17 — Lease worker dan penolakan hasil usang

Lease memberi worker kepemilikan sementara atas pekerjaan. Token generasi meningkat setiap kali lease baru diberikan; hasil worker lama ditolak meskipun akhirnya selesai.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
lease_lama=1 lease_baru=2
hasil_lama=ditolak
hasil_baru=diterima
status=selesai
```

## Telusuri program

Worker pertama memegang generasi 1. Setelah lease kedaluwarsa, worker kedua mendapat generasi 2. Pemeriksaan token menolak penyelesaian oleh pemegang lama; waktu kedaluwarsa juga membatasi hasil yang bisa diterima.

## Latihan mandiri

Coba selesaikan pekerjaan dengan token 1 pada waktu 4, sebelum lease kedaluwarsa. Hasilnya seharusnya diterima.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-16/README.md) · [Bab berikutnya](../bab-18/README.md)
