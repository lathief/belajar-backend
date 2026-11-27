# Bab 08 — Promise, async/await, dan kegagalan I/O

`readFile` versi promise mewakili pembacaan yang hasilnya tersedia kemudian. `await` menunggu hasil dalam fungsi `async`. `JSON.parse` mengubah teks menjadi data. Operasi baca dan parsing masing-masing bisa gagal, sehingga pemanggil membungkus keduanya dalam `try/catch`.

## Baca dan jalankan

Buka [main.mjs](main.mjs) dan [fixture pesanan](pesanan.json). Dari direktori `learn-javascript/bab-08`, jalankan perintah berikut.

```sh
node main.mjs
```

Keluaran program yang diharapkan

```text
PS-1 baru
```

## Telusuri program

Sukses membaca file menghasilkan teks. Setelah `JSON.parse`, barulah `kode` dan `status` dapat dibaca sebagai properti. File hilang dan JSON rusak mencapai `catch` melalui operasi yang berbeda.

## Latihan mandiri

Ubah status fixture menjadi `lunas` dan periksa output. Buat JSON rusak untuk memicu pesan kegagalan, lalu pulihkan contoh.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-07/README.md) · [Bab berikutnya](../bab-09/README.md)
