# Belajar PHP untuk backend developer

Anda akan beralih dari skrip terminal sederhana ke pemodelan pesanan, validasi input, dan simulasi pola backend. Contoh memakai PHP CLI dan pustaka standar; setiap bab dapat dijalankan mandiri tanpa Composer atau framework.

## Persiapan

Pasang PHP 8.2 atau lebih baru dan pastikan `php --version` menampilkan versinya. [Dokumentasi PHP](https://www.php.net/manual/en/)

Buka terminal di folder `learn-php`. Untuk mencoba bab pertama:

```sh
cd bab-01
php main.php
```

Jalankan setiap perintah dari folder bab yang disebutkan. Bab 09 memakai validator dari Bab 07 melalui path relatif; pertahankan susunan folder.

## Roadmap

| Bab | Materi | Yang dipraktikkan |
| --- | --- | --- |
| 01 | [Program pertama](bab-01/README.md) | Menjalankan skrip PHP dan membaca output |
| 02 | [Variabel, tipe, dan operator](bab-02/README.md) | Menghitung subtotal pesanan |
| 03 | [Kondisi dan perulangan](bab-03/README.md) | Menyaring jumlah yang memenuhi batas |
| 04 | [Fungsi dan nilai balik](bab-04/README.md) | Memisahkan aturan harga dari tampilan |
| 05 | [Array berindeks dan asosiatif](bab-05/README.md) | Membaca stok dan membedakan nilai yang hilang |
| 06 | [Class, property, dan method](bab-06/README.md) | Menjaga perubahan stok melalui method |
| 07 | [Interface dan aturan ongkir](bab-07/README.md) | Mengganti aturan ongkir lewat kontrak yang sama |
| 08 | [Parsing, exception, dan validasi](bab-08/README.md) | Membedakan input rusak dari pelanggaran aturan |
| 09 | [Proyek mini rekapan pesanan](bab-09/README.md) | Menyusun laporan dari argumen terminal |
| 10 | [Reservasi stok dan idempotensi](bab-10/README.md) | Mencegah pengiriman ulang mengurangi stok dua kali |
| 11 | [Transisi status pesanan](bab-11/README.md) | Menolak aksi yang tidak sesuai status |
| 12 | [Konflik penulisan dan nomor versi](bab-12/README.md) | Menolak perubahan dari formulir yang sudah usang |
| 13 | [Membangun ulang stok dari log peristiwa](bab-13/README.md) | Menangani peristiwa duplikat dan urutan yang terlewat |
| 14 | [Ekspor audit dengan paginasi cursor](bab-14/README.md) | Membaca halaman tanpa mengulang entri |
| 15 | [Token bucket untuk permintaan ekspor](bab-15/README.md) | Membedakan kapasitas lonjakan dan laju pengisian |
| 16 | [Retry dengan backoff dan deadline](bab-16/README.md) | Menghentikan percobaan sebelum anggaran waktu habis |
| 17 | [Lease worker dan penolakan hasil usang](bab-17/README.md) | Menolak hasil dari pemegang lease lama |
| 18 | [Proyek expert — outbox, inbox, dan pemulihan](bab-18/README.md) | Menguji efek bisnis di berbagai titik kegagalan |

## Cara berlatih

Baca penjelasan dan kode setiap bab. Prediksi output sebelum menjalankan program, lalu ubah input atau aturan sesuai latihan. Bab 09 memakai jumlah pesanan valid 1–5; uji input normal, batas, rusak, dan kosong. Bab 10–18 adalah model edukatif di memori, bukan pengganti transaksi dan penyimpanan persisten pada aplikasi produksi.
