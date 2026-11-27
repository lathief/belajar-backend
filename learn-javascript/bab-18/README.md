# Bab 18 — Proyek expert — outbox, inbox, dan pemulihan

**Tingkat expert** · Menguji efek bisnis di semua kombinasi tiga titik kegagalan.

## Studi kasus

Gudang menerima pesanan A untuk satu unit. Setelah stok berkurang, layanan pengemasan harus membuat satu paket. Respons klien bisa hilang, relay bisa terputus, dan pengakuan konsumen ke broker bisa hilang. Anda akan menguji sendiri pada batas mana pengulangan dapat menggandakan efek.

## Aturan kasus

1. **Produsen** mencatat pesanan, mengurangi stok awal 5, dan mencatat pesan tertunda di outbox dalam satu commit model. Percobaan ulang memakai identitas pesanan A yang sama.
2. **Relay** menerbitkan A ke antrean lalu menandai outbox terkirim. Jika putus di antaranya, relay menerbitkan ID yang sama lagi.
3. **Konsumen** merekam ID di inbox dan membuat paket dalam satu commit model. Pesan dihapus dari antrean setelah pengakuan (ack). Jika ack hilang, pesan di kepala antrean dibaca lagi.

## Baca dan jalankan

Buka [main.mjs](main.mjs). Dari direktori `learn-javascript/bab-18`, jalankan perintah berikut.

```sh
node main.mjs
```

Keluaran yang diharapkan

```text
stok=4 publikasi=2 paket=1
skenario_lulus=8
```

## Telusuri program

Mengganti objek produsen memodelkan satu commit. `Set` menyaring ID yang pernah diterima. `shift()` mengakui pesan sesudah efek lokal tercatat.

Program memulai keadaan baru untuk setiap kombinasi. Masing-masing gangguan disuntikkan paling banyak satu kali. Ada 2 × 2 × 2 = 8 lintasan. Saat ketiganya terjadi, stok tetap 4, relay menerbitkan dua salinan A, dan paket tetap satu. Pemeriksaan program juga menuntut tepat satu publikasi pada lintasan tanpa gangguan relay. Keadaan yang dianggap tersimpan dimodelkan di memori antarpercobaan. Pada sistem nyata, gunakan transaksi persisten untuk setiap commit lokal, ID pesan stabil, serta batas unik pesanan dan inbox. Pembuatan paket di sini berarti pembuatan record bisnis lokal.

## Latihan mandiri

1. Lakukan uji mutasi pada salinan contoh. Pindahkan penambahan paket ke luar pemeriksaan inbox, lalu jalankan delapan lintasan. Temukan kombinasi kegagalan paling sederhana yang membuat pemeriksaan gagal. Pulihkan penjagaan inbox dan pastikan semua lulus lagi.
2. Perluas model menjadi dua pesanan A dan B, masing-masing satu unit. Gunakan identitas per pesanan dan per pesan. Suntikkan kegagalan hanya untuk A. Dari stok awal 5, hasil akhir harus stok 3 dan dua record paket. Jelaskan tempat menyimpan kunci idempotensi bila pembuatan paket memanggil API eksternal.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-17/README.md)
