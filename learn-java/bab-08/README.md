# Bab 08 — Exception dan validasi

`Integer.parseInt` mengubah teks ke bilangan bulat atau melempar `NumberFormatException`. `try` membatasi operasi yang diawasi. `catch` menangani kegagalan konversi. Angka yang berhasil dibaca tetap perlu diperiksa terhadap aturan pesanan.

## Baca dan jalankan

Buka [Main.java](Main.java). Dari direktori `learn-java/bab-08`, jalankan perintah berikut.

```sh
java Main.java
```

Keluaran program yang diharapkan

```text
jumlah=2
di luar batas
bukan angka
```

## Telusuri program

Untuk `0`, parsing berhasil dan operator `kondisi ? a : b` memilih pesan batas. Untuk `x`, konversi gagal sebelum ekspresi itu dievaluasi.

## Latihan mandiri

Tambahkan `-1` dan `6`. Terapkan batas 1–5 sehingga keduanya menghasilkan `di luar batas` dan input berikutnya tetap diproses.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-07/README.md) · [Bab berikutnya](../bab-09/README.md)
