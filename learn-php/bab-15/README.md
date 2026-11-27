# Bab 15 — Token bucket untuk permintaan ekspor

Token bucket membatasi laju rata-rata sekaligus mengizinkan lonjakan singkat. Kapasitas menentukan jumlah maksimum token yang bisa ditampung.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
diizinkan
diizinkan
ditolak
diizinkan
token=1
```

## Telusuri program

Bucket mulai dengan dua token dan mengisi ulang satu token per detik hingga kapasitas dua. Permintaan ketiga terlalu cepat, sementara permintaan setelah dua detik dapat diterima; satu token tersisa.

## Latihan mandiri

Ubah kapasitas menjadi 3 dan prediksi hasil tiga permintaan awal.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-14/README.md) · [Bab berikutnya](../bab-16/README.md)
