# Belajar C# untuk backend developer

Anda akan memodelkan pesanan dengan property dan record, merangkum koleksi memakai LINQ, dan menangani input kosong atau rusak dengan nullable dan `TryParse`. Sembilan bab dasar ini dimulai dari program satu baris dan berakhir dengan alat terminal untuk merekap pesanan.

## Persiapan

Pasang .NET SDK 10 atau lebih baru untuk menjalankan file `.cs` secara langsung. Pastikan `dotnet --list-sdks` menampilkan SDK tersebut. [Dokumentasi C#](https://learn.microsoft.com/en-us/dotnet/csharp/)

Buka terminal di folder `learn-c-sharp`. Untuk mencoba bab pertama, jalankan perintah berikut.

```sh
cd bab-01
dotnet run --file Program.cs
```

## Roadmap

| Bab | Materi | Yang dipraktikkan |
| --- | --- | --- |
| 01 | [Program pertama dengan .NET](bab-01/README.md) | Menjalankan program dan membaca output |
| 02 | [Variabel, tipe, dan operator](bab-02/README.md) | Menghitung nilai pesanan |
| 03 | [Kondisi dan iterasi](bab-03/README.md) | Menyaring jumlah yang memenuhi batas |
| 04 | [Fungsi dan parameter](bab-04/README.md) | Memanggil aturan harga dari fungsi |
| 05 | [Class, property, dan objek](bab-05/README.md) | Membatasi perubahan stok melalui property |
| 06 | [Record, enum, dan switch](bab-06/README.md) | Menyalin record dan memetakan status |
| 07 | [Koleksi generik dan LINQ](bab-07/README.md) | Menghitung unit dengan LINQ |
| 08 | [TryParse, out, dan nullable](bab-08/README.md) | Memvalidasi teks dan menangani nilai null |
| 09 | [Proyek mini rekapan pesanan](bab-09/README.md) | Menyusun laporan dari argumen terminal |

## Tingkat intermediate

| Bab | Studi kasus | Yang dipraktikkan |
| --- | --- | --- |
| 10 | [Reservasi stok dan idempotensi](bab-10/README.md) | Menangani pengiriman ulang dan konflik ID tanpa mengurangi stok dua kali |
| 11 | [Transisi status pesanan](bab-11/README.md) | Menolak aksi yang tidak cocok dengan status saat ini |
| 12 | [Konflik penulisan dan nomor versi](bab-12/README.md) | Menjaga agar formulir lama tidak menimpa perubahan terbaru |
| 13 | [Membangun ulang stok dari log peristiwa](bab-13/README.md) | Menangani peristiwa duplikat dan nomor urut yang terlewat |
| 14 | [Ekspor audit dengan paginasi cursor](bab-14/README.md) | Membaca log per halaman sambil menerima entri baru |
| 15 | [Token bucket untuk permintaan ekspor](bab-15/README.md) | Membedakan kapasitas lonjakan dari laju pengisian |

## Tingkat advanced

| Bab | Studi kasus | Yang dipraktikkan |
| --- | --- | --- |
| 16 | [Retry dengan backoff dan deadline](bab-16/README.md) | Menghentikan percobaan sebelum anggaran waktu habis |
| 17 | [Lease worker dan penolakan hasil usang](bab-17/README.md) | Menguji batas waktu kepemilikan dan token generasi |

## Tingkat expert

| Bab | Studi kasus | Yang dipraktikkan |
| --- | --- | --- |
| 18 | [Proyek expert — outbox, inbox, dan pemulihan](bab-18/README.md) | Menguji efek bisnis di semua kombinasi tiga titik kegagalan |

## Cara berlatih

Baca penjelasan bab dan kode sumbernya. Prediksi output sebelum menjalankan program, lalu bandingkan hasilnya dengan contoh. Kerjakan latihan dengan mengubah kode dan jalankan kembali. Lanjutkan ketika Anda dapat menjelaskan mengapa perubahan input mengubah hasil.

Pada proyek dasar Bab 09, perlakukan setiap argumen sebagai input dari satu klien. Jumlah pesanan yang diterima berada di rentang 1–5. Uji input normal, input di batas, input rusak, dan daftar kosong. Kelima roadmap menggunakan kasus ini agar hasil dasar antarbahasa dapat dibandingkan.
