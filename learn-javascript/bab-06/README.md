# Bab 06 — Transformasi koleksi dan Map

Arrow function `n => ...` dapat diberikan sebagai callback. `filter` memilih elemen, `map` membentuk nilai baru, dan `reduce` menggabungkan hasil ke satu akumulator. `Map` menghubungkan label dengan hasil untuk pencarian selanjutnya.

## Baca dan jalankan

Buka [main.mjs](main.mjs). Dari direktori `learn-javascript/bab-06`, jalankan perintah berikut.

```sh
node main.mjs
```

Keluaran program yang diharapkan

```text
omzet=25000
```

## Telusuri program

Hasil antara `filter` ialah `[2, 3]` dan hasil `map` ialah `[10000, 15000]`. `reduce` memulai dari 0, menambah 10000, lalu 15000.

## Latihan mandiri

Tambahkan 5 dan perketat filter menjadi 1–5. Untuk harga 5000, ringkasan harus menjadi 50000. Coba array kosong dan jelaskan peran nilai awal 0.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-05/README.md) · [Bab berikutnya](../bab-07/README.md)
