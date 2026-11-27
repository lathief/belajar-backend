# Bab 07 — Interface dan aturan ongkir

Interface menetapkan kontrak yang dapat dipenuhi beberapa class. Pemanggil dapat mengganti strategi ongkir tanpa mengubah fungsi penghitung.

## Baca dan jalankan

```sh
php main.php
```

Keluaran yang diharapkan:

```text
total=59000
total=52000
```

## Telusuri program

Kedua strategi menerapkan `AturanOngkir`. Fungsi `hitungTagihan` menerima interface tersebut, bukan bergantung pada class konkret.

## Latihan mandiri

Buat strategi ongkir gratis dengan batas subtotal, lalu gunakan pada `hitungTagihan`.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-06/README.md) · [Bab berikutnya](../bab-08/README.md)
