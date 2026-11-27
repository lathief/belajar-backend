# Bab 07 — Koleksi generik dan LINQ

`List<int>` menyimpan data bertipe sama. LINQ menerima lambda seperti `n => n > 0`. `Where` memilih elemen dan `Sum` menghitung jumlahnya. `Dictionary` menyimpan hasil berdasarkan kunci teks.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-07`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs
```

Keluaran program yang diharapkan

```text
unit=5
```

## Telusuri program

`Where` membentuk urutan angka positif. `Sum()` membacanya sampai selesai dan mengembalikan 5. Nilai itu lalu disimpan di bawah kunci `unit`.

## Latihan mandiri

Tambahkan 8 dan 5. Ubah filter menjadi 1–5, sehingga nilai laporan menjadi 10. Coba juga daftar kosong dan amati hasil `Sum`.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-06/README.md) · [Bab berikutnya](../bab-08/README.md)
