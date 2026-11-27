# Bab 12 — Konflik penulisan dan nomor versi

## Studi kasus

Dua operator membuka alamat Bandung pada versi 1. A menyimpan perubahan ke Solo. B masih memegang formulir lama. Anda akan menggunakan nomor versi untuk mendeteksi konflik saat B menyimpan Bogor.

**Target belajar** — Menjaga agar formulir lama tidak menimpa perubahan terbaru.

## Aturan kasus

Salin versi saat membaca. Jika versi yang dikirim cocok dengan versi tersimpan, ubah kota dan naikkan versi satu kali. Jika berbeda, pertahankan kota dan versi. Untuk percobaan terakhir, B membaca ulang data terbaru dan pengguna menyetujui penggantiannya.

## Baca dan jalankan

Buka [main.rs](main.rs). Dari direktori `learn-rust/bab-12`, jalankan contoh berikut.

```sh
cargo run --quiet --bin bab-12
```

Keluaran yang diharapkan

```text
A tersimpan kota=Solo versi=2
B konflik kota=Solo versi=2
B tersimpan kota=Bogor versi=3
```

## Telusuri program

`&mut Alamat` memberi akses untuk memperbarui satu objek. Dua salinan `u32` merekam versi yang dilihat masing-masing operator.

Kedua pembacaan terjadi sebelum penulisan pertama sehingga konflik dapat diulang tanpa thread atau jeda waktu. Pada baris kedua, versi 2 milik perubahan A tetap tersimpan. Dalam database dengan penulis bersamaan, gabungkan pemeriksaan versi dan perubahan kota dalam satu operasi bersyarat, misalnya `UPDATE ... WHERE id = ... AND versi = ...` yang juga menaikkan versi. Pemanggil perlu menangani hasil nol baris terubah.

## Latihan mandiri

1. Tukar urutan penyimpanan pertama A dan B. Sesudah dua percobaan, B harus berhasil menyimpan Bogor pada versi 2 dan A mengalami konflik.
2. Kirim dua perubahan memakai versi baca yang sama. Buktikan hanya satu yang berhasil. Jelaskan mengapa pengiriman ulang dengan versi terbaru perlu mempertimbangkan maksud pengguna.

[Roadmap](../README.md) · [Bab sebelumnya](../bab-11/README.md) · [Bab berikutnya](../bab-13/README.md)
