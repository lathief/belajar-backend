# Bab 03 — Percabangan dan perulangan

Array `int[]` menyimpan beberapa angka. Bentuk `for (int n : daftar)` mengambil nilai berikutnya pada setiap putaran. `if` menjalankan blok hanya ketika perbandingan di kiri dan kanan `&&` sama-sama benar.

## Baca dan jalankan

Buka [Main.java](Main.java). Dari direktori `learn-java/bab-03`, jalankan perintah berikut.

```sh
java Main.java
```

Keluaran program yang diharapkan

```text
valid=2
```

## Telusuri program

Untuk 0, syarat `n >= 1` sudah salah. Untuk 8, syarat pertama benar tetapi `n <= 5` salah. Hanya 1 dan 4 yang mencapai `valid++`.

## Latihan mandiri

Tambahkan input 5 dan 6. Hasil harus menjadi 3. Cetak angka yang ditolak dengan menambahkan `else`.

Ganti enhanced `for` dengan `int i = 0;` dan `while (i < daftar.length)`. Baca `daftar[i]` sebagai nilai saat ini, lalu naikkan `i` di setiap iterasi. Periksa bahwa hasil dengan input awal tetap 2. Apa yang terjadi jika indeks berhenti bertambah?

[Roadmap](../README.md) · [Bab sebelumnya](../bab-02/README.md) · [Bab berikutnya](../bab-04/README.md)
