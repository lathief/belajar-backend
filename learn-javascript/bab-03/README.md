# Bab 03 — Kondisi, array, dan loop

`for...of` membaca nilai array secara berurutan. `if` memilih cabang berdasarkan kondisi, sedangkan `&&` meminta kedua perbandingan benar. Untuk membandingkan nilai dan tipe secara langsung, gunakan `===`.

## Baca dan jalankan

Buka [main.mjs](main.mjs). Dari direktori `learn-javascript/bab-03`, jalankan perintah berikut.

```sh
node main.mjs
```

Keluaran program yang diharapkan

```text
valid=2
```

## Telusuri program

Saat `n` bernilai 4, kedua perbandingan benar. Ketika `n` bernilai 8, hanya perbandingan pertama yang benar sehingga counter tetap 2.

## Latihan mandiri

Tambahkan 5 dan 6 sehingga hasil menjadi 3. Lalu coba string `"2"`. Tambahkan `Number.isInteger(n)` agar hanya nilai bertipe angka bulat yang diterima.

Coba bentuk `while (i < daftar.length)` dengan `let i = 0`. Baca nilai lewat `daftar[i]` lalu naikkan `i` di setiap iterasi. Periksa bahwa hasil dengan input awal tetap 2. Apa yang terjadi jika indeks berhenti bertambah?

[Roadmap](../README.md) · [Bab sebelumnya](../bab-02/README.md) · [Bab berikutnya](../bab-04/README.md)
