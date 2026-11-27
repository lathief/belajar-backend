# Bab 07 — List, Map, dan generics

`List<String>` menyimpan urutan kode, termasuk kode berulang. `Map<String, Integer>` menyimpan hitungan per kode. Parameter generik memberi compiler informasi mengenai jenis kunci dan nilai. `getOrDefault` menentukan nilai awal ketika kunci baru muncul.

## Baca dan jalankan

Buka [Main.java](Main.java). Dari direktori `learn-java/bab-07`, jalankan perintah berikut.

```sh
java Main.java
```

Keluaran program yang diharapkan

```text
PS-1=2
```

## Telusuri program

Kunci `PS-1` pertama kali dibaca dengan default 0 dan disimpan sebagai 1. Pada kemunculan kedua, map mengembalikan 1 lalu `put` menggantinya dengan 2.

## Latihan mandiri

Tambahkan `PS-2` dua kali. Ubah laporan agar menampilkan kedua hitungan dalam urutan PS-1 lalu PS-2. Hasilnya 2 dan 3.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-06/README.md) · [Bab berikutnya](../bab-08/README.md)
