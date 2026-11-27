# Bab 05 — Class, property, dan objek

Konstruktor menerima nilai awal ketika `new` membuat objek. Property dengan `get; private set;` dapat dibaca dari luar class, sedangkan perubahan melewati method class. Pada file ini, definisi tipe ditulis setelah pernyataan tingkat atas.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-05`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs
```

Keluaran program yang diharapkan

```text
sisa=1
```

## Telusuri program

Konstruktor berjalan ketika `new` dipanggil. Sintaks `=>` di konstruktor menulis isi yang hanya memiliki satu ekspresi. Membaca `stok.Sisa` dari luar class memanggil getter.

## Latihan mandiri

Panggil `AmbilSatu()` tiga kali untuk stok awal 2. Hasil akhir harus 0. Coba menulis `stok.Sisa = 9` dari luar class dan jelaskan pesan compiler.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-04/README.md) · [Bab berikutnya](../bab-06/README.md)
