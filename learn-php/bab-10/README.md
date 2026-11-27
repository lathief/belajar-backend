# Bab 10 — Reservasi stok dan idempotensi

Setiap permintaan logis memiliki ID. Gudang menyimpan keputusan pertama per ID agar percobaan ulang tidak mengurangi stok lagi. Input invalid tidak dicatat sehingga ID masih bisa dikoreksi.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
A 3 diterima stok=2
B 3 stok_kurang stok=2
C 0 jumlah_invalid stok=2
C 2 diterima stok=0
A 3 diterima stok=0
A 1 konflik stok=0
B 3 stok_kurang stok=0
```

## Telusuri program

Riwayat dicari sebelum stok diperiksa. ID yang sama dan jumlah yang sama mengembalikan keputusan awal; jumlah berbeda menghasilkan konflik.

## Latihan mandiri

Tambahkan pengiriman ulang `A 3` setelah konflik. Hasilnya tetap `diterima` dan stok tidak berubah.

## Batas model

Riwayat hanya hidup di memori proses. Aplikasi nyata memerlukan penyimpanan persisten dan transaksi yang atomik.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-09/README.md) · [Bab berikutnya](../bab-11/README.md)
