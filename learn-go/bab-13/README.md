# Bab 13 — Membangun ulang stok dari log peristiwa

## Studi kasus

Kartu stok dihitung dari riwayat barang masuk dan keluar. Seusai gangguan, peristiwa nomor 3 tiba sebelum nomor 2 dan nomor 2 dikirim dua kali. Anda akan membangun proyeksi, yaitu ringkasan yang bisa dihitung ulang dari log.

**Target belajar** — Menangani peristiwa duplikat dan nomor urut yang terlewat.

## Aturan kasus

Mulai dari posisi dan stok 0. Satu aliran gudang memakai nomor berurutan mulai 1. Setiap nomor selalu membawa delta yang sama. Nomor ≤ posisi sudah diproses. Nomor > posisi + 1 menandakan celah. Hanya nomor tepat berikutnya yang mengubah stok dan posisi.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-13`, jalankan contoh berikut.

```sh
go run main.go
```

Keluaran yang diharapkan

```text
1 diterapkan posisi=1 stok=5
3 celah posisi=1 stok=5
2 diterapkan posisi=2 stok=4
2 duplikat posisi=2 stok=4
3 diterapkan posisi=3 stok=2
```

## Telusuri program

Zero value `ringkasan` langsung menjadi titik awal replay. Dua `return` pada `switch` mencegah mutasi ketika nomor peristiwa belum memenuhi syarat.

Peristiwa 1 menambah lima unit. Ketika 3 tiba terlalu awal, posisi harus tetap 1 agar 2 bisa diterapkan. Salinan 2 diabaikan dan pengiriman ulang 3 akhirnya mengurangi dua unit. Sumber contoh mengirim ulang secara eksplisit. Untuk proyeksi persisten, simpan stok dan posisi terakhir dalam transaksi yang sama.

## Latihan mandiri

1. Putar `(1,+5), (2,-1), (3,-2)` dari ringkasan kosong. Putar lagi pada ringkasan yang sama. Kedua kali harus berakhir pada posisi 3 dengan stok 2.
2. Hilangkan semua kemunculan nomor 2 dari contoh. Posisi harus tertahan di 1 dan stok tetap 5. Jelaskan kapan konsumen perlu meminta pengiriman ulang dari sumber.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-12/README.md) · [Bab berikutnya](../bab-14/README.md)
