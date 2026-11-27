# Bab 04 — Fungsi dan parameter

Fungsi lokal dapat dideklarasikan di berkas dengan top-level statements. Tipe hasil mendahului nama fungsi. `return` mengirim nilai ke pemanggil. Parameter `decimal` dan `int` memisahkan harga dari banyaknya item.

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-04`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs
```

Keluaran program yang diharapkan

```text
total=15000
```

## Telusuri program

Pemanggilan mengikat `5000m` ke parameter `harga` dan 3 ke `jumlah`. Nilai `decimal` hasil fungsi langsung dipakai dalam string interpolasi.

## Latihan mandiri

Tambahkan parameter biaya kirim dengan nilai default `0m`. Pastikan pemanggilan awal tetap menghasilkan 15000 dan biaya 1000 menghasilkan 16000.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-03/README.md) · [Bab berikutnya](../bab-05/README.md)
