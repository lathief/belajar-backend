# Bab 08 — TryParse, out, dan nullable

`TryParse` mengembalikan `bool` dan mengisi variabel `out` dengan hasil konversi. `int?` menyatakan kemungkinan nilai tidak tersedia. Fungsi mengembalikan `null` untuk format salah maupun jumlah di luar batas.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-08`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs
```

Keluaran program yang diharapkan

```text
2
ditolak
ditolak
```

## Telusuri program

Untuk `x`, `TryParse` mengembalikan `false`. Untuk `0`, parsing sukses tetapi syarat rentang gagal. Keduanya menghasilkan `null`. Operator `?.` dan `??` menerjemahkannya ke teks `ditolak`.

## Latihan mandiri

Tambahkan `-1`, `5`, dan `6`. Prediksi mana yang diterima sebelum menjalankan. Hanya `5` yang menambah keluaran valid.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-07/README.md) · [Bab berikutnya](../bab-09/README.md)
