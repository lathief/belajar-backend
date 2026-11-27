# Bab 05 — Ownership, borrowing, dan teks

`String` memiliki teks yang ukurannya bisa berubah. Memindahkan `String` ke variabel lain memindahkan kepemilikannya. `&str` memberi fungsi akses baca ke teks tanpa mengambil pemiliknya. `trim()` menghasilkan potongan teks yang meminjam sumber.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-05`, jalankan perintah berikut.

```sh
cargo run --quiet --bin bab-05
```

Keluaran program yang diharapkan

```text
kode=PS-7
panjang asli=6
```

## Telusuri program

Literal awal memiliki enam byte ASCII termasuk dua spasi. `bersihkan(&kode)` meminjam teks dan menghasilkan potongan empat byte. Karena pemiliknya masih `kode`, baris berikutnya dapat membaca panjang aslinya.

## Latihan mandiri

Coba pindahkan `kode` ke `salinan` sebelum baris cetak dan amati pesan compiler saat `kode` dipakai lagi. Ganti pemindahan itu dengan peminjaman dan buat program berjalan kembali.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-04/README.md) · [Bab berikutnya](../bab-06/README.md)
