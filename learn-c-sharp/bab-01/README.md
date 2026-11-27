# Bab 01 — Program pertama dengan .NET

SDK .NET dapat menjalankan file C# secara langsung sejak .NET 10. Top-level statements menjadi titik masuk program. `Console.WriteLine` memanggil method dari pustaka standar untuk mencetak satu baris.

Direktif `#:property PublishAot=false` memberi tahu SDK untuk memakai runtime .NET lokal saat mengerjakan contoh ini. [Panduan aplikasi satu file](https://learn.microsoft.com/en-us/dotnet/core/sdk/file-based-apps)

## Baca dan jalankan

Buka [Program.cs](Program.cs). Dari direktori `learn-c-sharp/bab-01`, jalankan perintah berikut.

```sh
dotnet run --file Program.cs
```

Keluaran program yang diharapkan

```text
layanan pesanan siap
```

## Telusuri program

`using System` membuat nama `Console` tersedia. `WriteLine` menerima teks dalam tanda kurung. Compiler menyusun titik masuk dari pernyataan yang Anda tulis di luar class.

## Latihan mandiri

Tambahkan `mode lokal` sebagai baris kedua. Hilangkan titik koma sekali, amati lokasi kesalahan compiler, lalu perbaiki.

[Roadmap](../README.md) · [Bab berikutnya](../bab-02/README.md)
