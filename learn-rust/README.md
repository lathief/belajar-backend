# Belajar Rust untuk backend developer

Anda akan meminjam data untuk validasi, memodelkan status dengan enum, dan menangani input yang gagal dikonversi melalui `Result`. Sembilan bab dasar ini dimulai dari program satu baris dan berakhir dengan alat terminal untuk merekap pesanan.

## Persiapan

Pasang toolchain stable melalui rustup. Contoh menggunakan Cargo dan edition 2024 (Rust 1.85+). Periksa dengan `rustc --version` dan `cargo --version`. [Dokumentasi Rust](https://doc.rust-lang.org/book/) · [Dasar Pemrograman Rust oleh Noval Agung](https://dasarpemrogramanrust.novalagung.com/)

Buka terminal di folder `learn-rust`. Untuk mencoba bab pertama, jalankan perintah berikut.

```sh
cd bab-01
cargo run --quiet --bin bab-01
```

## Roadmap

| Bab | Materi | Yang dipraktikkan |
| --- | --- | --- |
| 01 | [Program pertama](bab-01/README.md) | Menjalankan program dan membaca output |
| 02 | [Nilai, tipe, dan operator](bab-02/README.md) | Menghitung nilai pesanan |
| 03 | [Kondisi dan perulangan](bab-03/README.md) | Menyaring jumlah yang memenuhi batas |
| 04 | [Fungsi dan nilai balik](bab-04/README.md) | Memanggil aturan harga dari fungsi |
| 05 | [Ownership, borrowing, dan teks](bab-05/README.md) | Meminjam teks tanpa memindahkan pemiliknya |
| 06 | [Struct, enum, dan match](bab-06/README.md) | Memetakan status pesanan |
| 07 | [Koleksi dan pencarian](bab-07/README.md) | Menyimpan item dan mencari stok |
| 08 | [Result, parsing, dan operator ?](bab-08/README.md) | Membedakan error parsing dan error domain |
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
