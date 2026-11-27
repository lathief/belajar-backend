# Bab 16 — Retry dengan backoff dan deadline

**Tingkat advanced** · Menghentikan percobaan sebelum anggaran waktu habis.

## Studi kasus

Setelah mendapat izin memulai ekspor, layanan meminta satu halaman data. Layanan tujuan sesekali menjawab 503. Anda akan memutuskan kapan mencoba lagi dan kapan pemanggil harus menerima hasil gagal.

## Aturan kasus

Kasus ini membaca data, sehingga aman diulang. Setiap daftar berisi minimal satu respons dan panjangnya sekaligus menentukan batas percobaan. Respons 200 menyelesaikan operasi. Hanya 503 yang dapat dicoba kembali. Mulai pada t=0 dengan deadline positif dan jeda awal 1 detik. Gandakan jeda setelah retry dijadwalkan. Percobaan baru hanya boleh mulai jika waktunya **kurang dari** deadline.

## Baca dan jalankan

Buka [main.mjs](main.mjs). Dari direktori `learn-javascript/bab-16`, jalankan perintah berikut.

```sh
node main.mjs
```

Keluaran yang diharapkan

```text
t=0 percobaan=1 kode=503
t=1 percobaan=2 kode=503
t=3 percobaan=3 kode=200
hasil=sukses
t=0 percobaan=1 kode=503
t=1 percobaan=2 kode=503
hasil=deadline
t=0 percobaan=1 kode=400
hasil=permanen
```

## Telusuri program

`entries()` menghasilkan indeks dan kode respons. Waktu berupa angka virtual sehingga perilaku deadline tidak bergantung pada penjadwalan event loop.

Dua kasus pertama memakai respons identik. Jadwalnya 0, 1, dan 3. Deadline 4 memberi kesempatan ketiga. Deadline 3 menghentikannya tepat di batas. Respons 400 menghentikan kasus terakhir walaupun masih ada respons berikutnya. Jam virtual membuat ketiga lintasan dapat diulang tanpa menunggu. Durasi panggilan diasumsikan nol. Dalam I/O nyata, batasi juga waktu panggilan aktif dengan sisa deadline.

## Latihan mandiri

1. Ganti daftar menjadi `[503, 503]` dengan deadline 10. Harapkan tepat dua percobaan dan `hasil=habis`. Coba `[200]` dan pastikan hanya ada satu panggilan.
2. Tambahkan durasi satu detik untuk setiap panggilan sebelum menghitung jadwal retry. Untuk deadline 4 dan respons `[503, 503, 200]`, hitung kapan panggilan kedua selesai dan mengapa panggilan ketiga tidak boleh dimulai.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-15/README.md) · [Bab berikutnya](../bab-17/README.md)
