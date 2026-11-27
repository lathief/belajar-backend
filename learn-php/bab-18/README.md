# Bab 18 — Proyek expert — outbox, inbox, dan pemulihan

**Tingkat expert** · Menguji efek bisnis pada tiga titik kegagalan.

## Studi kasus

Pesanan mengurangi stok dan mencatat pesan outbox. Relay dapat putus setelah publish; konsumen dapat kehilangan ack setelah membuat paket. Setiap skenario menguji bahwa pengulangan tidak menggandakan efek bisnis.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
stok=4 publikasi=2 paket=1
skenario_lulus=8
```

## Telusuri program

Tiga kegagalan dinyalakan atau dimatikan secara independen, menghasilkan 2 × 2 × 2 lintasan. ID stabil dan inbox mencegah duplikasi; model memverifikasi stok tetap berkurang sekali dan satu paket dibuat.

## Latihan mandiri

Hapus pemeriksaan inbox sebelum membuat paket, lalu jalankan lagi. Perhatikan invariant yang gagal; pulihkan pemeriksaan tersebut.

## Batas model

State disimpan di memori untuk latihan. Aplikasi nyata memerlukan transaksi persisten untuk pesanan/outbox dan inbox/efek bisnis.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-17/README.md)
