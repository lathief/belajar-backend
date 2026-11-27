# Bab 05 — Class, objek, dan enkapsulasi

`new` membuat objek dan memanggil konstruktor. `this` menunjuk objek yang sedang digunakan. Field `private` membatasi akses langsung dari luar class. Dua objek dapat menjalankan method yang sama dengan state berbeda.

## Baca dan jalankan

Buka [Main.java](Main.java). Dari direktori `learn-java/bab-05`, jalankan perintah berikut.

```sh
java Main.java
```

Keluaran program yang diharapkan

```text
sisa=1
```

## Telusuri program

Konstruktor menerima 2, lalu `this.sisa` menyimpan nilai di objek. Setelah `ambilSatu()` dipanggil, `sisa()` membaca field objek yang sama dan menghasilkan 1.

## Latihan mandiri

Buat dua objek dengan stok awal 2 dan 4. Ambil satu dari objek pertama. Cetak sisa keduanya dan pastikan hasilnya 1 dan 4.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-04/README.md) · [Bab berikutnya](../bab-06/README.md)
