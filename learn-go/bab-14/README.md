# Bab 14 — Ekspor audit dengan paginasi cursor

## Studi kasus

Petugas audit mengekspor dua entri per halaman. Sementara itu, sistem masih menambahkan entri. Anda akan memakai ID terakhir yang terkirim sebagai cursor dan ID terbesar saat sesi dimulai sebagai batas atas ekspor.

**Target belajar** — Membaca log per halaman sambil menerima entri baru.

## Aturan kasus

Log hanya bertambah di akhir. ID positif, unik, terurut naik, dan boleh memiliki celah. Ukuran halaman minimal 1. Tangkap batas atas sekali pada awal sesi. Ambil hingga sejumlah ukuran halaman dengan `cursor < id <= batas`. Pindah ke ID terakhir yang terkirim. Halaman kosong mengakhiri sesi.

## Baca dan jalankan

Buka [main.go](main.go). Dari direktori `learn-go/bab-14`, jalankan contoh berikut.

```sh
go run main.go
```

Keluaran yang diharapkan

```text
halaman=101,104
halaman=107,112
susulan=115
```

## Telusuri program

Slice keluaran menampung ID yang lolos. Saat `append` menambah log, salinan batas yang bertipe `int` tetap sama.

Halaman pertama berhenti pada 104. Entri 115 datang sebelum halaman kedua, tetapi batas ekspor tetap 112. Pembacaan susulan melanjutkan cursor terakhir dengan batas atas baru. Contoh memindai koleksi dalam memori. Untuk data besar, indeks pada ID memungkinkan pencarian dimulai langsung setelah cursor.

## Latihan mandiri

1. Ubah ukuran menjadi 3. Ekspor pertama harus mencetak `101,104,107` lalu `112`. Entri 115 tetap muncul pada pembacaan susulan.
2. Uji fungsi pembaca dengan cursor 112 dan batas 112, lalu dengan daftar kosong. Keduanya harus menghasilkan halaman kosong. Jelaskan apa yang dapat terjadi jika penulis terus menambah entri dan batas atas selalu dihitung ulang.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-13/README.md) · [Bab berikutnya](../bab-15/README.md)
