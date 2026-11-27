# Belajar Go untuk backend developer

Dalam sembilan bab dasar, Anda akan beralih dari mencetak log startup ke membuat CLI yang memvalidasi dan merekap pesanan. Kasus pesanan yang sama juga dipakai di roadmap Rust, Java, JavaScript, dan C# sehingga perilaku program mudah dibandingkan.

## Persiapan

Pasang Go dan periksa instalasinya dengan `go version`. Semua bab memakai pustaka standar dan dapat dijalankan langsung dari satu file. [Panduan instalasi Go](https://go.dev/doc/install)

Dari folder `learn-go`, jalankan contoh pertama.

```sh
cd bab-01
go run main.go
```

## Roadmap

| Bab | Materi | Yang dipraktikkan |
| --- | --- | --- |
| 01 | [Program pertama](bab-01/README.md) | Menulis pesan startup ke terminal |
| 02 | [Variabel, tipe, dan operator](bab-02/README.md) | Menghitung subtotal dalam rupiah |
| 03 | [Kondisi dan perulangan](bab-03/README.md) | Menyaring jumlah yang memenuhi batas |
| 04 | [Fungsi dan nilai balik](bab-04/README.md) | Memisahkan aturan harga dari tampilan |
| 05 | [Slice dan map](bab-05/README.md) | Membedakan stok nol dari kode yang hilang |
| 06 | [Struct, method, dan pointer](bab-06/README.md) | Mengurangi stok melalui method |
| 07 | [Interface dan aturan ongkir](bab-07/README.md) | Mengganti aturan ongkir lewat kontrak yang sama |
| 08 | [Parsing, error, dan validasi](bab-08/README.md) | Membedakan kesalahan format dan batas bisnis |
| 09 | [Proyek mini rekapan pesanan](bab-09/README.md) | Mengolah argumen terminal menjadi laporan |

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

Buka kode contoh, prediksi hasilnya, lalu jalankan dan bandingkan dengan keluaran di README bab. Kerjakan latihan dengan mengubah input atau aturan. Pada proyek dasar Bab 09, coba satu batch campuran, satu batch di batas rentang, dan satu batch kosong. Setelah memahami dasar ini, Anda juga bisa membaca [contoh worker](worker/README.md) yang tersedia di folder ini.

## Contoh server HTTP

Setiap contoh tersedia dalam versi Fiber, Gin, dan `net/http`.

- [API gateway](apigateway/README.md) meneruskan permintaan produk ke backend lokal.
- [Rate limiter](ratelimiter/README.md) membatasi jumlah permintaan ke satu endpoint.
- [Basic auth](basicauth/README.md) melindungi endpoint profil dengan nama pengguna dan kata sandi.
