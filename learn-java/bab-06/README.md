# Bab 06 — Record, enum, dan switch

`record` mendefinisikan pembawa data dengan accessor seperti `kode()`. `enum` membatasi daftar status. Ekspresi `switch` menghasilkan satu nilai, sehingga label baru dapat ditambahkan di satu tempat.

## Baca dan jalankan

Buka [Main.java](Main.java). Dari direktori `learn-java/bab-06`, jalankan perintah berikut.

```sh
java Main.java
```

Keluaran program yang diharapkan

```text
PS-1 lunas
```

## Telusuri program

`var` membuat compiler menyimpulkan tipe lokal `Pesanan`. Accessor `status()` menghasilkan `Status.BAYAR`, sehingga ekspresi `switch` menghasilkan teks `lunas`.

## Latihan mandiri

Tambahkan status `BATAL` dan label `dibatalkan`. Buat satu record untuk membuktikan cabang tersebut dipakai.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-05/README.md) · [Bab berikutnya](../bab-07/README.md)
