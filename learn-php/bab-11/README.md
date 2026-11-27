# Bab 11 — Transisi status pesanan

Pesanan hanya boleh berpindah status melalui transisi yang diizinkan. Method domain menolak aksi yang tidak cocok dengan status saat ini.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
baru -> dibayar
ditolak: pesanan_harus_dibayar
dibayar -> dikirim
```

## Telusuri program

Pengiriman sebelum pembayaran ditolak dan status tetap `baru`. Setelah pembayaran, pesanan dapat dikirim.

## Latihan mandiri

Coba membayar pesanan yang sudah dikirim. Tambahkan aturan agar aksi tersebut ditolak.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-10/README.md) · [Bab berikutnya](../bab-12/README.md)
